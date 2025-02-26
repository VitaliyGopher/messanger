package server

import (
	"net/http"

	"github.com/VitaliyGopher/messanger/internal/pkg/model"
	_ "github.com/VitaliyGopher/messanger/internal/pkg/swagger"
	"github.com/gin-gonic/gin"
)

// Ping			godoc
// @Summary 	Ping-pong
// @Description	ping-pong))
// @Produce 	application/json
// @Tags		ping
// @Success 	200
// @Router 		/ping [get]
func (s *server) Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
}

// SendVerificationCode 			godoc
// @Summary 						Returns a verification code
// @Description 					Returns a verification code for auth
// @Param 							email body swagger.SendCodeParams{} true "email"
// @Produce							application/json
// @Tags 							auth
// @Success 						200
// @Success 						400
// @Success 						500
// @Router							/sendCode [post]
func (s *server) SendCodeHandler(c *gin.Context) {
	type request struct {
		Email string `json:"email"`
	}
	var req request
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

// GetJWT			godoc
// @Summary			Returns jwt tokens
// @Description		Returns jwt tokens and checks verification code
// @Param			email_and_code body swagger.GetJWTParams{} true "email and verification code"
// @Produce 		application/json
// @Tags			auth
// @Success 		201
// @Success 		400
// @Success 		500
// @Router			/createJwt [post]
func (s *server) GetJWT(c *gin.Context) {
	type request struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}
	var req request

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

// GetNewJWT		godoc
// @Summary			Returns jwt tokens
// @Description		Refresh jwt tokens
// @Param			refresh_token body swagger.GetNewJWTParams{} true "refresh token"
// @Produce 		application/json
// @Tags			auth
// @Success 		200
// @Success 		400
// @Router			/getNewJWT [post]
func (s *server) GetNewJWT(c *gin.Context) {
	type request struct {
		Refresh string `json:"refresh_token"`
	}
	var req request

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
