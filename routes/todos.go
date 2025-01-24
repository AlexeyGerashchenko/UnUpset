package routes

import (
	"github.com/gin-gonic/gin"
)

// Дописать

func SetupRotesTodos(r *gin.Engine) {
	todos := r.Group("/todos")
	{
		todos.GET("/")
		todos.GET("/:user_id")
		todos.POST("/:user_id")
		todos.PATCH("/:user_id/:task_id")
		todos.DELETE("/:user_id/:task_id")
	}
}
