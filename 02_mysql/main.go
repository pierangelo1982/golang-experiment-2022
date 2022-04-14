package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type city struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func dbConnection() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var dbUser string = os.Getenv("DATABASE_USERNAME")
	var dbPassword string = os.Getenv("DATABASE_PASSWORD")
	var dbHost string = os.Getenv("DATABASE_HOST")
	var dbPort string = os.Getenv("DATABASE_PORT")

	fmt.Println(dbUser, dbPassword, dbHost, dbPort)
	var dbUrl string = fmt.Sprintf("%s:%s@tcp(%s:%s)/demo", dbUser, dbPassword, dbHost, dbPort)
	fmt.Println("@@@@@@@@@", dbUrl)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		//panic(err)
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func getCities(c *gin.Context) {
	db, error := dbConnection()
	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "connessione database fallita"})
		return
	}

	rows, err := db.Query("SELECT id, name, country FROM City;")
	if err != nil {
		//panic(err)
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "connessione database fallita"})
		return
	}

	var cities []*city

	for rows.Next() {
		var rowCity city
		err := rows.Scan(&rowCity.ID, &rowCity.Name, &rowCity.Country)
		if err != nil {
		}

		cities = append(cities, &rowCity)
	}
	err = rows.Err()

	defer db.Close()

	c.IndentedJSON(http.StatusOK, cities)
}

func main() {
	router := gin.Default()
	router.GET("/", getCities)
	router.Run("localhost:8080")
}
