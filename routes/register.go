package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/models"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSONP(400, gin.H{"message": "Failed parse current id, please try again"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		c.JSONP(500, gin.H{"message": "Failed fetch event with current id, please try again"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		c.JSONP(500, gin.H{"message": "Could not register user with the events, please try again"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Event registered!",
	})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSONP(400, gin.H{"message": "Failed parse current id, please try again"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		c.JSONP(500, gin.H{"message": "Could not cancel event registration, please try again"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Event canceleld!",
	})
}
