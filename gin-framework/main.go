package main

import (
	"github.com/Emrul-Hasan-Emon/demo-go-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load the Templates
	router.LoadHTMLGlob("templates/*")

	// Define the route handler
	router.GET("/", handler.ShowHomePage)
	router.GET("/article/view/:article_id", handler.GetSingleArtilce)
	router.Run()
}
