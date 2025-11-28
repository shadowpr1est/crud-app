package main

import (
	_ "crud-app/docs"
	"crud-app/internal/handler"
	"log"
)

// @title Task API
// @version 1.0
// @description CRUD API for learning Go.
// @host localhost:8080
// @BasePath /

func main() {
	router := handler.SetupRouter()

	log.Println("Server starting at 'http://localhost:8080'")
	log.Println("Swagger documentation available at 'http://localhost:8080/swagger/index.html'")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
