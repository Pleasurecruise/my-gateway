package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pleasure1234/my-gateway/common"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (h *HelloController) Hello(c *gin.Context) {
	name := c.DefaultQuery("name", "World")
	common.OK(c, gin.H{"message": "Hello, " + name + "!"})
}