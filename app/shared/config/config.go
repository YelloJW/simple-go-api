package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Env                string
	DbConnectionString string
}

func New() (*config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &config{Env: os.Getenv("ENV"), DbConnectionString: os.Getenv("DB_CONNECTION_STRING")}, nil
}
