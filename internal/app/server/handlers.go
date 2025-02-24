package server

import (
	"net/http"
	"strconv"

	"github.com/VitaliyGopher/messanger/internal/pkg/model"
	"github.com/gin-gonic/gin"
)

func (s *server) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}

func (s *server) SendCodeHandler(c *gin.Context) {
	email := c.PostForm("email")
	sms, err := s.verifyCode.SendCode(email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"code":       sms.Code,
		"smsExpires": sms.Timestamp,
	})
}

func (s *server) GetJWT(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")

	code_int, _ := strconv.Atoi(code)

	s.jwt.UserRepo.Create(&model.User{
		Email: email,
	})

	u, err := s.verifyCode.CheckCode(email, code_int)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, err := s.jwt.CreateAccessJWT(int(u.ID))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	refresh_token, err := s.jwt.CreateRefreshJWT(int(u.ID))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"access_token": access_token,
		"refresh_token": refresh_token,
	})
}
