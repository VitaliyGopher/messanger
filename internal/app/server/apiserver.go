package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	jwttoken "github.com/VitaliyGopher/messanger/internal/app/auth"
	"github.com/VitaliyGopher/messanger/internal/app/verification_code"
	"github.com/VitaliyGopher/messanger/internal/pkg/postgres"
	rsa_key "github.com/VitaliyGopher/messanger/pkg/rsa"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type server struct {
	router     *gin.Engine
	store      postgres.Storage
	verifyCode VerifyCodeInterface
	jwt        jwttoken.JWT
}

func newServer(store postgres.Storage, verifyCode VerifyCodeInterface, jwt jwttoken.JWT) *server {
	s := &server{
		router:     gin.Default(),
		store:      store,
		verifyCode: verifyCode,
		jwt:        jwt,
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
	verifyCodeRepo := postgres.NewVerifyCodeRepo(store)
	userRepo := postgres.NewUserRepo(store)

	verifyCode := verification_code.New(verifyCodeRepo, userRepo)

	privateKey, err := rsa_key.LoadOrGenerateRSA(os.Getenv("rsa_filename"))
	if err != nil {
		log.Fatal("rsa error: ", err)
	}

	jwt := jwttoken.New(privateKey, userRepo)

	s := newServer(*store, verifyCode, jwt)

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
