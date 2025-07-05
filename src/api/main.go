package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/test", test)

	err := router.Run()
	if err != nil {
		panic(fmt.Errorf("error starting server: %v", err))
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
