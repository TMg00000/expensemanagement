package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MONGO_URI    string `envconfig:"MONGO_URI" required:"true"`
	EXPENSES_COL string `envconfig:"EXPENSES_COL" required:"true"`
}

var Env Config

func StartConfig() error {
	_ = godotenv.Load()

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
