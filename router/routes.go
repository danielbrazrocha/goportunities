package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	print("chamei")
	//definindo rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
