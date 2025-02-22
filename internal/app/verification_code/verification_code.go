package verification_code

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/VitaliyGopher/messanger/internal/pkg/model"
	"github.com/VitaliyGopher/messanger/internal/pkg/postgres"
)

type VerifyCode struct { 
	VerifyCodeRepo  *postgres.VerifyCodeRepo
	UserRepo *postgres.UserRepo
}

func New(VerifyCodeRepo *postgres.VerifyCodeRepo, userRepo *postgres.UserRepo) *VerifyCode {
	return &VerifyCode{
		VerifyCodeRepo:  VerifyCodeRepo,
		UserRepo: userRepo,
	}
}

func (c *VerifyCode) SendCode(email string) (*model.VerifyCode, error) {
	verifyCode := &model.VerifyCode{
		Email: email,
	}

	verifyCode, err := c.VerifyCodeRepo.FindVerifyCode(verifyCode)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}

	if verifyCode.Code != 0 {
		timeDiference := verifyCode.Timestamp - time.Now().Unix()
		if (timeDiference) > 0 {
			return verifyCode, nil
		}

		c.VerifyCodeRepo.DeleteCode(verifyCode)
	}

	code := rand.Intn(9000) + 1000
	timestamp := time.Now().Add(5 * time.Minute).Unix()

	verifyCode = &model.VerifyCode{
		Email: email,
		Code:  code,
		Timestamp: timestamp,
	}

	c.VerifyCodeRepo.CreateCode(verifyCode)

	return verifyCode, nil
}

func (c *VerifyCode) CheckCode(email string, code int) (*model.User, error) {
	check_code := &model.VerifyCode{
		Email: email,
		Code: code,
	}

	db_code, err := c.VerifyCodeRepo.FindVerifyCode(check_code)
	if err != nil {
		return nil, err
	}

	if check_code.Code != db_code.Code {
		return nil, errors.New("wrong code")
	}

	if db_code.Timestamp < time.Now().Unix() {
		return nil, errors.New("sms is expired")
	}

	u, err := c.UserRepo.FindByEmail(check_code.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not exists")
		}
		return nil, err
	}

	return u, nil
}