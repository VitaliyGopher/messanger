package auth

import (
	"crypto/rsa"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessJWT(uid int, privateKey *rsa.PrivateKey) (string, error) {
	exp, _ := strconv.Atoi(os.Getenv("jwt_access_exp"))
	claims := jwt.MapClaims{
		"sub": uid,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(exp) * time.Minute).Unix(),
		"type": "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshJWT(uid int, privateKey *rsa.PrivateKey) (string, error) {
	exp, _ := strconv.Atoi(os.Getenv("jwt_refresh_exp"))
	claims := jwt.MapClaims{
		"sub": uid,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(exp) * time.Hour * 24).Unix(),
		"type": "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string, publicKey *rsa.PublicKey) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid signing method")
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}
