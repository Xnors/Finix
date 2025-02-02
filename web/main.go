package main

import (
	"finix-web/consts"
	"finix-web/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/create", routes.RCreate)
	r.Run(":" + consts.PORT)
}
