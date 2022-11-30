package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"mymodule14-html-form-with-gin/data"
	"net/http"
)

func postUserRegistration(c *gin.Context) {

	userRegistration := &data.UserRegistration{}

	if err := c.ShouldBind(userRegistration); err != nil {
		verrs := err.(validator.ValidationErrors)
		// ---- convert validation errors to list of error messages (strings)
		messages := make([]string, len(verrs))
		log.Printf("verrs: %#v", len(verrs))
		for i, verr := range verrs {
			messages[i] = fmt.Sprintf("%s: %s", verr.Field() /*verr.Error()*/, "The field is required")
		}
		c.HTML(http.StatusBadRequest, "form.html.tmpl", gin.H{
			"errors":           messages,
			"userRegistration": userRegistration,
		})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/user-created")
}

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
	engine.POST("/form", postUserRegistration)

	// ---- after successful form submission the user gets redirected to here
	engine.POST("/user-created", func(c *gin.Context) {
		c.HTML(http.StatusCreated, "user-created.html.tmpl", gin.H{
			"appName": "Fancy App",
		})
	})

	engine.Run(":9090")
}
