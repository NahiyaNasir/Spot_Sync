package config
import (
    "log"
    "os"

    "github.com/joho/godotenv"
)
type Config struct {
	Port string `json:"port"`
	Dsn string `json:"dsn"`
}
func LoadEnv()	(*Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Dsn: os.Getenv("DSN"),
	}
}