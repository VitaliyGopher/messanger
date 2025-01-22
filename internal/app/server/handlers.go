package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *server) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}

func (s *server) SendSmsCodeHandler(c *gin.Context) {
	phone := c.PostForm("phone")
	sms, err := s.sms.SendSmsCode(phone)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"code":       sms.Code,
		"smsExpires": sms.Timestamp - time.Now().Unix(),
	})
}
