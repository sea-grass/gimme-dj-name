package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GimmeDJName(c *gin.Context) {
	name := c.Param("name")
	dj_name := "xx_[]**++_" + name + "_++**[]_xx"
	c.String(http.StatusOK, dj_name)
}

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/xx_gimme_dat_dj_name/:name", GimmeDJName)

	router.Run(port)
}
