package router

import (
	"github.com/VitaliyGopher/messanger/internal/handlers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func init() {
	router.GET("/ping", handlers.Ping)
}

func NewRouter() *gin.Engine {
	return router
}
