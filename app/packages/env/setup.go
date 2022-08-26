package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	Env string
}

func Load() environment {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	return environment{Env: os.Getenv("ENV")}
}