package server

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VitaliyGopher/messanger/internal/app/sms"
	"github.com/VitaliyGopher/messanger/internal/pkg/postgres"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type server struct {
	router *gin.Engine
	store  postgres.Storage
	sms    SmsInterface
}

func newServer(store postgres.Storage, sms SmsInterface) *server {
	s := &server{
		router: gin.Default(),
		store:  store,
		sms: sms,
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

	s := newServer(*store, sms)

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
