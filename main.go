package main

import (
	"net/http"

	"example.com/test/db"
	"example.com/test/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

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

	//? this is going to fetch the data from the incomming request and put it into the variable of event
	err := context.ShouldBindJSON(&event)

	//? this is going to check weather the incomming data is comming in way that was said
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}

	//? manually assign the value for UserId and eventID
	event.UserId = 1
	event.ID = 1

	event.Save()
	//? I am return what I have save
	context.JSON(http.StatusCreated, gin.H{"this is the data of event that you have saved": event})
}
