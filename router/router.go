package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default() //inicialize GinGonic Router com default

	initializeRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
