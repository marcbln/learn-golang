package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	const bindAddress = ":9000"

	router := httprouter.New()
	router.GET("/hello", getRoot)

	server := http.Server{
		Addr:    bindAddress,
		Handler: router,
	}

	log.Info().Msg("starting server on " + bindAddress)
	err := server.ListenAndServe()
	log.Fatal().Err(err).Msg("server failed")
}

func getRoot(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello World"))
}
