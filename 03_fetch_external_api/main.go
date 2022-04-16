package main

import (
	"fetch-external-api/nhlApi"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	// help benchmarking request time
	now := time.Now()

	rosterFile, err := os.OpenFile("roster.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("errore nell'aprire il file roster.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatal("errore nell'ottenere i teams: %v", err)
	}

	for _, team := range teams {
		log.Println("---------------------------")
		log.Printf("Name %s", team.Name)

		log.Println("---------------------------")
	}

	// tempo impiegato
	log.Printf("took %v", time.Now().Sub(now))
}
