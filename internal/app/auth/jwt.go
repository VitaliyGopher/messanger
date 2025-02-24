package jwttoken

import (
	"crypto/rsa"
	"errors"

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
		UserRepo:   userRepo,
	}
}

func (t *JWT) CreateAccessJWT(uid int) (string, error) {
	token, err := auth.GenerateAccessJWT(uid, t.privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *JWT) CreateRefreshJWT(uid int) (string, error) {
	token, err := auth.GenerateRefreshJWT(uid, t.privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *JWT) GetAccessJWT(refresh string) (string, error) {
	claims, err := auth.VerifyJWT(refresh, &t.privateKey.PublicKey)
	if err != nil {
		return "", err
	}

	if claims["type"] != "refresh" {
		return "", errors.New("is not refresh token")
	}

	token, err := t.CreateAccessJWT(int(claims["sub"].(float64)))
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

	if claims["type"] != "access" {
		return nil, errors.New("is not access token")
	}

	u, err := t.UserRepo.FindByID(claims["iss"].(uint))
	if err != nil {
		return nil, err
	}

	return u, nil
}
