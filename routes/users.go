package routes

import (
	"UnUpset/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRotesUsers(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", controllers.GetAllUsers)
		users.GET("/:id")
		users.POST("/", controllers.CreateUser)
		users.PATCH("/:id")
		users.DELETE("/:id")
	}
}
