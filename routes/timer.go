package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRotesTimer(r *gin.Engine) {
	timer := r.Group("/timer")
	{
		timer.GET("/")
		timer.GET("/:id")
		timer.PATCH("/:id")
	}
}
