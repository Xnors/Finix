package main

import (
	"finix-web/consts"
	"finix-web/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Finix!"})
	})
	r.GET("/create", routes.RCreate)
	r.GET("/delete", routes.RDelete)
	r.Run(":" + consts.PORT)
}
