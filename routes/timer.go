package routes

import (
	"UnUpset/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutesTimer(r *gin.Engine) {
	timer := r.Group("/timer")
	{
		timer.GET("/", controllers.GetAllTimers)
		timer.GET("/:id", controllers.GetTimersById)
		timer.PATCH("/:id", controllers.ChangeTimerById)
	}
}
