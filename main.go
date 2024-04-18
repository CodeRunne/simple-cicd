package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", PongController)
	fmt.Println("server started at port :8004")
	router.RunTLS(":8004", "server.crt", "server.key")
}

func PongController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "pong",
	})
}
