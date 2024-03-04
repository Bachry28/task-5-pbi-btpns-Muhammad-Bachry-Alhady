package main

import (
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/database"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/helpers"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/router/photorouter"
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/router/userrouter"
	"github.com/gin-gonic/gin"
)

// Init initializes the application
func Init() {
	// Load environment variables
	helpers.LoadEnv()
	// Connect to the database
	database.ConnectDatabase()

}

// Run starts the application server
func Run() {
	// Initialize the Gin router
	r := gin.Default()

	// Set up routes for user-related endpoints
	userrouter.UserRouter(r)

	// Set up routes for photo-related endpoints
	photorouter.PhotoRouter(r)

	// Start the Gin server
	r.Run()
}

func main() {
	// Call the Init function to initialize the application
	Init()

	// Call the Run function to start the application server
	Run()
}
