package helpers

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.New()
	return router
}