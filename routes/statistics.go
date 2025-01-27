package routes

import (
	"UnUpset/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutesStatistics(r *gin.Engine) {
	statistics := r.Group("/statistics")
	{
		statistics.GET("/", controllers.GetAllStatistics)
		statistics.GET("/:id", controllers.GetStatisticsById)
		statistics.PATCH("/:id", controllers.ChangeStatisticsById)
	}
}
