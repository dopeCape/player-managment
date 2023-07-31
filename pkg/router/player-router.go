package router

import (
	controller "github.com/dopeCape/player-managment/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var CreateRouter = func(r *gin.Engine) {
	r.POST("/players", controller.HandleNewPlayer)
	r.PUT("/players/:id", controller.HandleUpadtePlayer)
	r.DELETE("/players/:id", controller.HandleDeletePlayer)
	r.GET("/players", controller.HandelGetAllPlayer)
	r.GET("/players/rank/:val", controller.HandleGetPlayerWithRank)
	r.GET("/players/random", controller.HandelGetRandomPlayer)
}
