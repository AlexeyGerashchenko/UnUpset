package routes


import (
	"github.com/gin-gonic/gin"
)

func SetupRotesStatistics(r *gin.Engine) {
	statistics := r.Group("/statistics") {
		router.GET("/", )
		router.GET("/:id", )
		router.PATCH("/:id", )
	}
}
