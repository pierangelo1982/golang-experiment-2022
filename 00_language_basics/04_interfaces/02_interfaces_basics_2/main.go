package main

import (
	"fmt"
	"time"
)

type personAge interface {
	getPersonAge() int
	getisAdult() bool
}

type Person struct {
	Name      string
	BirthDate time.Time
}

/* func (p Person) getPersonAge() float64 {
	today := time.Now()
	diff := today.Sub(p.BirthDate)
	return (diff.Hours() / 24) / 365
} */

func (p Person) getPersonAge() int {
	today := time.Now()
	diff := today.Sub(p.BirthDate)
	return int((diff.Hours() / 24) / 365)
}

func (p Person) isAdult() bool {
	if p.getPersonAge() >= 18 {
		return true
	} else {
		return false
	}
}

func main() {
	input := "1982-01-09"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)

	fmt.Println(t.Format("2006-01-02"))

	var person Person
	person.Name = "John"
	person.BirthDate = t

	fmt.Println(person.getPersonAge())
	fmt.Println(person.isAdult())

}
