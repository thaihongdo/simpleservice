package router

import (
	"simpleservice/cmd/entity-server/api"

	"github.com/gin-gonic/gin"
)

func initRouter(Router *gin.RouterGroup) {
	wagerRouter := Router.Group("wagers")
	{
		wagerRouter.POST("", api.AddWager)
		wagerRouter.GET("", api.GetWager)
	}

	buyRouter := Router.Group("buy")
	{
		buyRouter.POST("/:wager-id", api.Buy)
	}
}
