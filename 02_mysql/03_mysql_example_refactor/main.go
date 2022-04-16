package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"first-mysql/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Drivers:", sql.Drivers())
	dbConnection := utils.ReturnDatabaseUrlFromEnvVar()
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
