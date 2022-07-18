package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter(env string) *gin.Engine {
	var router = gin.Default()
	apiGroup := router.Group(fmt.Sprintf("%s", env))
	initRouter(apiGroup)
	return router
}
