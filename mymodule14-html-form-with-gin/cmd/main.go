package main

import (
	"github.com/gin-gonic/gin"
	"mymodule14-html-form-with-gin/data"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	// ---- display html form
	engine.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html.tmpl", gin.H{
			"appName": "Fancy App",
		})
	})

	// ---- handle html form
	engine.POST("/handle-form", func(c *gin.Context) {
		userRegistration := &data.UserRegistration{}
		if err := c.Bind(userRegistration); err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		c.String(http.StatusOK, "%#v", userRegistration)
	})

	engine.Run(":9090")
}
