package router

import (
	"portobe/pkg/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// For Generaral Routes

	{
		router.GET("/healt", controller.Health)
	}

	return router
}
