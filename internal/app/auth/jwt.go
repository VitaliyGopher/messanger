package jwttoken

import (
	"crypto/rsa"

	"github.com/VitaliyGopher/messanger/internal/pkg/auth"
	"github.com/VitaliyGopher/messanger/internal/pkg/model"
	"github.com/VitaliyGopher/messanger/internal/pkg/postgres"
)

type JWT struct {
	privateKey *rsa.PrivateKey
	UserRepo   *postgres.UserRepo
}

func New(privateKey *rsa.PrivateKey, userRepo *postgres.UserRepo) JWT {
	return JWT{
		privateKey: privateKey,
		UserRepo: userRepo,
	}
}

func (t *JWT) CreateJWT(uid int) (string, error) {
	token, err := auth.GenerateJWT(uid, t.privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *JWT) ParseToken(tokenStr string) (*model.User, error) {
	claims, err := auth.VerifyJWT(tokenStr, &t.privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	u, err := t.UserRepo.FindByID(claims["iss"].(uint))
	if err != nil {
		return nil, err
	}

	return u, nil
}
