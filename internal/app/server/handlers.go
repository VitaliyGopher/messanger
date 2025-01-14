package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}