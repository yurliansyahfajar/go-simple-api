package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	//Events Routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//Users Routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
