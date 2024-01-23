package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//?Define the routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getOneEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteOneEvent)
}
