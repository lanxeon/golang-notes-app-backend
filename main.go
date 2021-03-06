package main

import (
	"log"

	config "notes_app/config"
	middlewares "notes_app/middlewares"
	routes "notes_app/routes"

	gin "github.com/gin-gonic/gin"
)

func main() {

	// connect to db
	config.Connect()

	// Init Router
	router := gin.Default()

	// CORS middleware
	router.Use(middlewares.Cors)

	// Route Handlers / Endpoints
	routes.Routes(router)
	log.Fatal(router.Run(":4747"))
}
