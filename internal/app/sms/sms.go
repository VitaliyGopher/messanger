package sms

import (
	"math/rand"
	"time"

	"github.com/VitaliyGopher/messanger/internal/pkg/model"
	"github.com/VitaliyGopher/messanger/internal/pkg/postgres"
)

type Sms struct {
	SmsRepo  *postgres.SmsRepo
	UserRepo *postgres.UserRepo
}

func New(smsRepo *postgres.SmsRepo, userRepo *postgres.UserRepo) *Sms {
	return &Sms{
		SmsRepo:  smsRepo,
		UserRepo: userRepo,
	}
}

func (s *Sms) SendSmsCode(phone string) (*model.Sms, error) {
	sms := &model.Sms{
		Phone: phone,
	}

	sms, err := s.SmsRepo.FindSmsCode(sms)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
	}

	if sms.Code != 0 {
		timeDiference := sms.Timestamp - time.Now().Unix()
		if (timeDiference) > 0 {
			return sms, nil
		}

		s.SmsRepo.DeleteSmsCode(sms)
	}

	code := rand.Intn(9000) + 1000
	timestamp := time.Now().Add(5 * time.Minute).Unix()

	sms = &model.Sms{
		Phone: phone,
		Code:  code,
		Timestamp: timestamp,
	}

	s.SmsRepo.CreateSmsCode(sms)

	return sms, nil
}
