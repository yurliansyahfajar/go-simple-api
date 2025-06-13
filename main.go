package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "events list",
	})
}
