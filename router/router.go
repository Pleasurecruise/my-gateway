package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pleasure1234/my-gateway/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	registerAPIRoutes(r)

	return r
}