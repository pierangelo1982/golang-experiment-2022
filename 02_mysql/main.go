package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var dbUser string = os.Getenv("DATABASE_USERNAME")
var dbPassword string = os.Getenv("DATABASE_PASSWORD")
var dbHost string = os.Getenv("DATABASE_HOST")
var dbPort string = os.Getenv("DATABASE_PORT")

type city struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func dbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:alnitek82@tcp(0.0.0.0:3306)/demo")
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
