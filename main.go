package main

import (
	"example.com/test/db"
	"example.com/test/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	//? Initialize the server
	server := gin.Default()

	//here we pass the pointer of the server gin.Default automatically returns the pointer
	routes.RegisterRoutes(server)

	//? Run the server
	server.Run(":8080")
}
