package routes

import (
	"UnUpset/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutesTodos(r *gin.Engine) {
	todos := r.Group("/todos")
	{
		todos.GET("/", controllers.GetAllTodos)
		todos.GET("/:user_id", controllers.GetTodosByUserId)
		todos.POST("/:user_id", controllers.CreateTodos)
		todos.PATCH("/:user_id/:task_id", controllers.ChangeTodosById)
		todos.DELETE("/:user_id/:task_id", controllers.DeleteTodosById)
	}
}
