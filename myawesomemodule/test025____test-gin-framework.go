package main

import (
	"github.com/nicholasjackson/env"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	log.Info().Msgf("listening on %s", *bindAddress)

	engine.Run(*bindAddress) // listen and serve on localhost:8080
}
