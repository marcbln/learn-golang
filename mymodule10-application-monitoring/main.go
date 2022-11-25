package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"mymodule10-application-monitoring/handlers"
	"mymodule10-application-monitoring/metrics"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	const bindAddress = ":9000"

	mc := metrics.NewMetricsCollection(time.Now())
	router := httprouter.New()
	router.GET("/", handlers.GetRoot("hello world!", mc))
	router.GET("/health", handlers.GetHealth(mc))

	server := http.Server{
		Addr:    bindAddress,
		Handler: router,
	}

	log.Info().Msg("starting server on " + bindAddress)
	err := server.ListenAndServe()
	log.Fatal().Err(err).Msg("server failed")
}
