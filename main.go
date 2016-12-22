package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FirstHalf(s string) string {
	r := NewReader(s)
	len := Len(r)
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
//	router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hi")
		//c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/xx_gimme_dat_dj_name/*name", func(c *gin.Context) {
		name = gin.DefaultQuery("name", "")
		if name == "" {
			message := "Gotta gimme dat name doe"
		}
		message := "xx_**[]//_" + FirstHalf(name) + "xx" + SecondHalf(name) + "_\\[]**_xx"

		c.String(http.StatusOK, message)
	})

	router.Run(":" + port)
}
