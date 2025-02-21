package sms

import (
	"database/sql"
	"errors"
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
		if err != sql.ErrNoRows {
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

func (s *Sms) CheckSmsCode(phone string, code int) (*model.User, error) {
	check_sms := &model.Sms{
		Phone: phone,
		Code: code,
	}

	db_sms, err := s.SmsRepo.FindSmsCode(check_sms)
	if err != nil {
		return nil, err
	}

	if check_sms.Code != db_sms.Code {
		return nil, errors.New("wrong code")
	}

	if db_sms.Timestamp < time.Now().Unix() {
		return nil, errors.New("sms is expired")
	}

	u, err := s.UserRepo.FindByPhoneNumber(check_sms.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not exists")
		}
		return nil, err
	}

	return u, nil
}