package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	BotToken    string `env:"BOT_TOKEN,notEmpty"`
	PostgresURL string `env:"DATABASE_URL,notEmpty"`
	SecretKey string `env:"SECRET_KEY,notEmpty"`

}


func (s *Config) Load() {
	godotenv.Load()
	if err := env.Parse(s); err != nil {
		log.Fatal("couldn't load config: %s", err.Error())
	}
}