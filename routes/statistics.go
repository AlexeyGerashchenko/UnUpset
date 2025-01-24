package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRotesStatistics(r *gin.Engine) {
	statistics := r.Group("/statistics")
	{
		statistics.GET("/")
		statistics.GET("/:id")
		statistics.PATCH("/:id")
	}
}
