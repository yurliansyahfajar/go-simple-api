package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(200, events)
}

func createEvent(c *gin.Context) {
	var event models.Event

	// read from request.body
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	c.JSON(201, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}
