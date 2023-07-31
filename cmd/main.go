package main

import (
	"github.com/dopeCape/player-managment/pkg/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	router.CreateRouter(r)
	r.Run(":8000")
}
