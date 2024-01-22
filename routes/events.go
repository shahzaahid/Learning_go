//? this one is responsible to handle all the response handlers
// ? this function is used to get the event

package routes

import (
	"net/http"
	"strconv"

	"example.com/test/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getOneEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not convert into the int"})
		return
	}
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
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

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event try again later"})
		return
	}
	//? I am return what I have save
	context.JSON(http.StatusCreated, gin.H{"this is the data of event that you have saved": event})
}
