package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(taskHandler *TaskHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/tasks/", taskHandler.GetAll)
	router.GET("/tasks/:id", taskHandler.GetByID)
	router.POST("/tasks/", taskHandler.Create)
	router.PUT("/tasks/:id", taskHandler.Update)
	router.DELETE("/tasks/:id", taskHandler.Delete)

	return router
}
