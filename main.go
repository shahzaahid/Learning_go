package main

import (
	"net/http"

	"example.com/test/models"
	"github.com/gin-gonic/gin"
)

func main() {

	//? Initialize the server
	server := gin.Default()

	//?Define the routes
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	//? Run the server
	server.Run(":8080")
}

// ? this function is used to get the event
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

// ? this function is used to recive the add and create the event
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}
	models.UserId = 1
	models.ID = 1
}
