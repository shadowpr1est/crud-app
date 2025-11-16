package main

import (
	_ "crud-app/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

// @title Task API
// @version 1.0
// @description Simple CRUD API for learning Go.
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	storage := NewTaskStorage()
	handler := newHandler(storage)

	r.GET("/tasks", handler.GetAll)
	r.GET("/tasks/:id", handler.GetByID)
	r.POST("/tasks", handler.Create)
	r.PUT("/tasks/:id", handler.Update)
	r.DELETE("/tasks/:id", handler.Delete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	r.Run(":8080")
}
