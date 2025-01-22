package routes

import (
"github.com/gin-gonic/gin"
)

func SetupRotesTimer (r *gin.Engine) {
	timer := r.Group("/timer") {
		router.GET("/", )
		router.GET("/:id", )
		router.PATCH("/:id", )
	}
}

