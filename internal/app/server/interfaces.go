package server

import "github.com/VitaliyGopher/messanger/internal/pkg/model"

type SmsInterface interface {
	SendSmsCode(phone string) (*model.Sms, error)
	CheckSmsCode(phone string, code int) (*model.User, error)
}
