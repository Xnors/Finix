package main

import (
	"finix-web/consts"
	"finix-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			200, 
			gin.H{
				"message": "Welcome to Finix!", 
				"advice": "You can try to visit '/create?table_name=your_table_name' to create a new table.",
			},
		)
	})
	r.GET("/create", routes.RCreate)
	r.GET("/delete", routes.RDelete)
	r.GET("/get", routes.RGet)
	r.GET("/get_all", routes.RGetAll)
	r.GET("/get_info", routes.RGetInfo)
	r.Run(":" + consts.PORT)
}
