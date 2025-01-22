package routes

import (
	"github.com/gin-gonic/gin"
	"UnUpset/controllers"
)

func SetupRotesUsers (r *gin.Engine) {
	users := r.Group("/users") {
		router.GET("/", controllers.GetAllUsers)
		router.GET("/:id", )
		router.POST("/", controllers.CreateUser)
		router.PATCH("/:id", )
		router.DELETE("/:id", )
	}
}
