package server

import (
	"crypto/rand"
	"crypto/rsa"
	"database/sql"
	"fmt"
	"log"

	jwttoken "github.com/VitaliyGopher/messanger/internal/app/auth"
	"github.com/VitaliyGopher/messanger/internal/app/sms"
	"github.com/VitaliyGopher/messanger/internal/pkg/postgres"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type server struct {
	router *gin.Engine
	store  postgres.Storage
	sms    SmsInterface
	jwt    jwttoken.JWT
}

func newServer(store postgres.Storage, sms SmsInterface, jwt jwttoken.JWT) *server {
	s := &server{
		router: gin.Default(),
		store:  store,
		sms:    sms,
		jwt:    jwt,
	}

	s.configureRouter()

	return s
}

func Start(config *Config) error {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DB_username,
		config.DB_password,
		config.Host,
		"5432",
		config.DB_name,
		config.DB_sslmode,
	)

	db, err := newDB(connStr)
	if err != nil {
		log.Fatal("storage error: ", err)
	}
	defer db.Close()

	store := postgres.New(db)
	smsRepo := postgres.NewSmsRepo(store)
	userRepo := postgres.NewUserRepo(store)

	sms := sms.New(smsRepo, userRepo)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("error in generating rsa keys: %s", err)
	}
	jwt := jwttoken.New(privateKey, userRepo)

	s := newServer(*store, sms, jwt)

	return s.router.Run(config.Host + config.Port)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
