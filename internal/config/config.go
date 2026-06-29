package config
import (
    "log"
    "os"

    "github.com/joho/godotenv"
)
type Config struct {
	Port string `json:"port"`
	Dsn string `json:"dsn"`
	jwtSecret string 
}
func LoadEnv()	(*Config) {
	err := godotenv.Load()
	if err != nil {
		 log.Println("No .env file found, using environment variables")
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Dsn: os.Getenv("DSN"),
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}