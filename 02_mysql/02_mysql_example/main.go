package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Drivers:", sql.Drivers())
	dbConnection := returnDatabaseConnectionUrlFromEnvVariable()
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		log.Fatal("Impossibile aprire una connessione con il database")
	}
	defer db.Close()
	results, err := db.Query("SELECT id, name, country FROM City")
	if err != nil {
		log.Fatal("Errore quando emetto query su tabella:", err)
	}
	defer results.Close()
	for results.Next() {
		var (
			id      int
			name    string
			country string
		)
		err = results.Scan(&id, &name, &country)
		if err != nil {
			log.Fatal("Non posso parse row", err)
		}
		fmt.Printf("id è %d, %s è in %s \n", id, name, country)
	}

	fmt.Println("-----------------------------------------")

	var (
		id      int
		name    string
		country string
	)
	err = db.QueryRow("SELECT id, name, country FROM City WHERE id = 1").Scan(&id, &name, &country)
	if err != nil {
		log.Fatal("Errore quando emetto query su tabella:", err)
	}
	fmt.Printf("id è %d, %s è in %s \n", id, name, country)

	fmt.Println("----------------------------------------")

	cities := []struct {
		name      string
		country   string
		updatedAt time.Time
	}{
		{"berlino", "germania", time.Now()},
		{"Madrid", "Spagna", time.Now()},
		{"Lisbona", "Portogallo", time.Now()},
	}

	statement, err := db.Prepare("INSERT INTO City (name, country, updatedAt) VALUES (?, ?, ?)")
	defer statement.Close()
	if err != nil {
		log.Fatal("Impossibile preparare lo statement", err)
	}
	for _, city := range cities {
		_, err = statement.Exec(city.name, city.country, city.updatedAt)
		if err != nil {
			log.Fatal("Impossibile eseguire lo statement: ", err)
		}
	}
}

func returnDatabaseConnectionUrlFromEnvVariable() (url string) {
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
