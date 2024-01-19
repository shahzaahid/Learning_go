package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//? Initialize the server
	server := gin.Default()

	//?Define the routes
	server.GET("/events", getEvents)
	server.GET("/joy", getJoy)

	//? Run the server
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "this is the response from the getEvents"})
}

func getJoy(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This is the respokse from the get joy function "})
}
