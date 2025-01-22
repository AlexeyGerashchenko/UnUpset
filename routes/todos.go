package routes

import (
	"github.com/gin-gonic/gin"
)

// Дописать 

func SetupRotesTodos(r *gin.Engine) {
	todos := r.Group("/todos") {
		router.GET("/", )
		router.GET("/:user_id", )
		router.POST("/:user_id", )
		router.PATCH("/:user_id/:task_id", )
		router.DELETE("/:user_id/:task_id", )
	}
}