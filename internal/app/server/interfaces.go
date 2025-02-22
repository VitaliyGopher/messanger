package server

import "github.com/VitaliyGopher/messanger/internal/pkg/model"

type VerifyCodeInterface interface {
	SendCode(email string) (*model.VerifyCode, error)
	CheckCode(email string, code int) (*model.User, error)
}
