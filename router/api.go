package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pleasure1234/my-gateway/controller"
)

func registerAPIRoutes(r *gin.Engine) {
	helloCtrl := controller.NewHelloController()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": true})
	})

	// v1 API group
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", helloCtrl.Hello)
	}
}