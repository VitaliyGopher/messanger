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

func (s *server) SendSmsCodeHandler(c *gin.Context) {
	phone := c.PostForm("phone")
	sms, err := s.sms.SendSmsCode(phone)
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
	phone := c.PostForm("phone")
	code := c.PostForm("code")

	code_int, _ := strconv.Atoi(code)

	s.jwt.UserRepo.Create(&model.User{
		PhoneNumber: phone,
	})

	u, err := s.sms.CheckSmsCode(phone, code_int)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.jwt.CreateJWT(int(u.ID))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"jwt": token})
}
