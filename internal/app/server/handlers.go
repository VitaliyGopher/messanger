package server

import (
	"net/http"

	"github.com/VitaliyGopher/messanger/internal/pkg/model"
	"github.com/gin-gonic/gin"
)

func (s *server) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}

func (s *server) SendCodeHandler(c *gin.Context) {
	type requset struct {
		Email string `json:"email"`
	}
	var req requset
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sms, err := s.verifyCode.SendCode(req.Email)
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
	type requset struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}
	var req requset

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.jwt.UserRepo.Create(&model.User{
		Email: req.Email,
	})

	u, err := s.verifyCode.CheckCode(req.Email, req.Code)
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
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}

func (s *server) GetNewJWT(c *gin.Context) {
	type requset struct {
		Refresh string `json:"refresh_token"`
	}
	var req requset

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := s.jwt.GetNewJWT(req.Refresh)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"access_token":  tokens["access"],
		"refresh_token": tokens["refresh"],
	})
}
