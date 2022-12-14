package main

import (
	"context"
	"github.com/nicholasjackson/env"
	"log"
	"mymodule7-rest-v2/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	err := env.Parse()
	if err != nil {
		log.Fatalln("error parsing env: ", err)
	}

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// ---- create a new serve mux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/products/", handlers.NewProductsHandler(l))
	serveMux.Handle("/hello", handlers.NewHelloHandler(l))
	serveMux.Handle("/goodbye", handlers.NewGoodbyeHandler(l))

	// ---- create a new server
	server := http.Server{
		Addr:         *bindAddress,
		Handler:      serveMux,
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
			//l.Printf("Error starting server: %s\n", err)
			//os.Exit(1)
			l.Fatal(err)
		}
	}()

	// ---- trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// ---- block until message is consumed
	sig := <-sigChan
	l.Println("received terminate - graceful shutdown", sig)

	// ---- shutdown with a deadline context
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
