package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRotesUsers (r *gin.Engine) {
	users := r.Group("/users") {
		router.GET("/", )
		router.GET("/:id", )
		router.POST("/", )
		router.PATCH("/:id", )
		router.DELETE("/:id", )
	}
}
