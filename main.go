package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FirstHalf(s string) string {
	//dos tuff..
	return s
}
func SecondHalf(s string) string {
	return ""
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hi")
	})
	router.GET("/xx_gimme_dat_dj_name/", func(c *gin.Context) {
		c.String(http.StatusOK, "gotta give me dat name doe")
	})
	router.GET("/xx_gimme_dat_dj_name/:name", func(c *gin.Context) {
		name := c.Param("name")
		message = "xx_**[]//_" + FirstHalf(name) + "xx" + SecondHalf(name) + "_\\[]**_xx"
		c.String(http.StatusOK, message)
	})

	router.Run(":" + port)
}
