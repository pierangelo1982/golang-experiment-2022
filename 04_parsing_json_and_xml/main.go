package main

import (
	"log"
	"parsing-json-xml/feeds"
	"parsing-json-xml/itunes"
)

func main() {
	ias := itunes.NewItunesApiServices()

	res, err := ias.Search("Flash F1 - Formula One Podcast")
	if err != nil {
		log.Fatalf("errore nella ricerca %s", err)
	}

	for _, s := range res.Results {
		log.Print("-------------------")
		log.Printf("Artist %s", s.ArtistName)
		log.Printf("Podcast Name %s", s.TrackName)
		log.Printf("Feed Url %s", s.FeedURL)

		feed, err := feeds.GetFeed(s.FeedURL)
		if err != nil {
			log.Fatalf("errore")
		}

		for _, pod := range feed.Channel.Item {
			log.Print("-------------------")
			log.Printf("Title %s", pod.Title)
			log.Printf("Duration %s", pod.Duration)
			log.Printf("Description %s", pod.Description)
			log.Printf("UR %sL", pod.Enclosure.URL)
		}
		log.Print("--------------------------------")
	}
}
