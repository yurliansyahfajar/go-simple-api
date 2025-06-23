package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/models"
	"github.com/yurliansyahfajar/go-simple-api/utils"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed fetch all events data, please try again."})
		return
	}
	c.JSON(200, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSONP(400, gin.H{"message": "Failed fetch event with current id, please try again"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed fetch events, please try again."})
		return
	}

	c.JSON(200, event)
}

func createEvent(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(401, gin.H{"message": "unauthorize user"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(401, gin.H{"message": "unauthorize user"})
		return
	}

	var event models.Event
	// read from request.body
	err = c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed create events, please try again."})
		return
	}

	c.JSON(201, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSONP(400, gin.H{"message": "Failed fetch event with current id, please try again"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed fetch event, please try again."})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(400, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed update event, please try again."})
		return
	}

	c.JSON(200, gin.H{
		"message": "Event updated!",
		"event":   updatedEvent,
	})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSONP(400, gin.H{"message": "Failed fetch event with current id, please try again"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed fetch event, please try again."})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		c.JSON(500, gin.H{"message": "Failed delete event, please try again."})
		return
	}

	c.JSON(200, gin.H{
		"message": "Event deleted!",
	})
}
