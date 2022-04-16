package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReturnDatabaseUrlFromEnvVar() (url string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var dbUser string = os.Getenv("DATABASE_USERNAME")
	var dbPassword string = os.Getenv("DATABASE_PASSWORD")
	var dbHost string = os.Getenv("DATABASE_HOST")
	var dbPort string = os.Getenv("DATABASE_PORT")

	var dbUrl string = fmt.Sprintf("%s:%s@tcp(%s:%s)/demo", dbUser, dbPassword, dbHost, dbPort)
	return dbUrl
}
