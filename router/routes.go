package router

import "github.com/gin-gonic/gin"
import "net/http"

func initializeRoutes(router *gin.Engine) {
	// definindo rota OLD style
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// definindo rota NEW style
	v1 := router.Group("/api/v1"); {
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
	}
}
