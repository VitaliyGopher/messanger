package rsa_key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func LoadOrGenerateRSA(filename string) (*rsa.PrivateKey, error) {
	privateKey, err := loadOrGeneratePrivateKey(filename)
	if err != nil {
		return nil, err
	}

	

	return privateKey, nil
}

func loadOrGeneratePrivateKey(filename string) (*rsa.PrivateKey, error) {
	privateKey, err := loadPrivateKeyFromFile(filename)
	if err == nil {
		return privateKey, nil
	}

	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	err = savePrivateKeyToFile(filename, privateKey)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func savePrivateKeyToFile(filename string, privateKey *rsa.PrivateKey) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, privateKeyPEM)
	if err != nil {
		return err
	}

	return nil
}
