package handler

import (
	"crud-app/internal/auth"
	"crud-app/internal/repository"
	"crud-app/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	taskRepo := repository.NewTaskMemoryRepository()
	userRepo := repository.NewUserMemoryRepository()

	jwtManager := auth.NewJWTManager("secret", 15*time.Minute)

	taskService := service.NewTaskService(taskRepo)
	authService := service.NewAuthService(userRepo, jwtManager)

	taskHandler := NewTaskHandler(taskService)
	authHandler := NewAuthHandler(authService)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)

	api := router.Group("/api", jwtManager.GinMiddleware())
	{
		api.GET("/tasks/", taskHandler.GetAll)
		api.GET("/tasks/:id", taskHandler.GetByID)
		api.POST("/tasks/", taskHandler.Create)
		api.PUT("/tasks/:id", taskHandler.Update)
		api.DELETE("/tasks/:id", taskHandler.Delete)
	}
	return router
}
