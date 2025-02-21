package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Load() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("cannot to load enviroment: %w", err)
	}

	return nil
}
