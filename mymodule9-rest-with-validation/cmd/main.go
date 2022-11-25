package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"log"
	"mymodule9-rest-with-validation/handlers"
	"mymodule9-rest-with-validation/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
var shutdownWait = env.Int("SHUTDOWN_WAIT", false, 15, "the duration in seconds for which the server gracefully shutdownWait for existing connections to finish")

func main() {
	err := env.Parse()
	if err != nil {
		log.Fatalln("error parsing env: ", err)
	}

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// ---- handlers
	productHandler := handlers.NewProductsHandler(l)

	// ---- create a new serve mux and register the handlers
	muxRouter := mux.NewRouter() // was http.ServeMux before

	// ---- middleware
	muxRouter.Use(middleware.LoggingMiddleware)

	// ---- router for GET requests
	getRouter := muxRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)
	// ---- router for PUT requests
	putRouter := muxRouter.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(middleware.ProductDeserializerMiddleware)
	// ---- router for POST requests
	postRouter := muxRouter.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProduct)
	postRouter.Use(middleware.ProductDeserializerMiddleware)

	// ---- create a new server
	server := http.Server{
		Addr:         *bindAddress,
		Handler:      muxRouter,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	//l.Println("starting server on", bindAddress)
	//err := server.ListenAndServe()
	//if err != nil {
	//	l.Println("Error starting server: %server\n", err)
	//	os.Exit(1)
	//}

	// ---- start the server (in a go func so that it's not gonna block)
	go func() {
		l.Printf("starting server on %s\n", *bindAddress)
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(77)
			// l.Fatal(err)
		}
	}()

	// ---- trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// ---- block until message is consumed
	sig := <-sigChan
	l.Println("received terminate - graceful shutdown", sig)

	// Create a deadline to shutdownWait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*shutdownWait)*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise shutdownWait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should shutdownWait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
