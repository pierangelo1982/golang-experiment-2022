package main

import (
	"fetch-external-api/nhlApi"
	"io"
	"log"
	"os"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(teams))

	results := make(chan []nhlApi.Roster)

	for _, team := range teams {
		go func(team nhlApi.Team) {
			roster, err := nhlApi.GetRosters(team.ID)
			if err != nil {
				log.Fatal("errore nell'ottenere i roster: %v", err)
			}

			results <- roster

			wg.Done()
		}(team)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	log.Println(results)
	display(results)

	// tempo impiegato
	log.Printf("took %v", time.Now().Sub(now))
}

func display(results chan []nhlApi.Roster) {
	for r := range results {
		for _, ros := range r {
			log.Println("---------------------------")
			log.Printf("ID %d\n", ros.Person.ID)
			log.Printf("Name %s\n", ros.Person.FullName)
			log.Printf("Position %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey %s\n", ros.JerseyNumber)
			log.Println("---------------------------")
		}

	}
}
