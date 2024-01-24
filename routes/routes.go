package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//?Define the routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getOneEvent)
	server.POST("/events", createEvent)          //auth
	server.PUT("/events/:id", updateEvent)       //auth
	server.DELETE("/events/:id", deleteOneEvent) //auth
	server.POST("/signup", signup)
	server.POST("/login", login)
}
