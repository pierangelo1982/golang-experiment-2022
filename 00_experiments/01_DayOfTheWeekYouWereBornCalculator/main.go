package main

import (
	"fmt"
	"log"
	"strconv"
)

type MonthNumberValue struct {
	Name   string
	Number int
}

var months []MonthNumberValue = []MonthNumberValue{
	{Name: "Gennaio", Number: 0},
	{Name: "Febbraio", Number: 3},
	{Name: "Marzo", Number: 3},
	{Name: "Aprile", Number: 6},
	{Name: "Maggio", Number: 1},
	{Name: "Giugno", Number: 4},
	{Name: "Luglio", Number: 6},
	{Name: "Agosto", Number: 2},
	{Name: "Settembre", Number: 5},
	{Name: "Ottobre", Number: 0},
	{Name: "Novembre", Number: 3},
	{Name: "Dicembre", Number: 5},
}

func getMonths() (int, error) {
	var selectedMonth int
	for index, month := range months {
		fmt.Printf("%d - %s\n", index, month.Name)
	}
	fmt.Print("Seleziona il mese in cui sei nato: ")
	_, err := fmt.Scanf("%d", &selectedMonth)
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	return months[selectedMonth].Number, err
}

func getDayNumberOfTheMonth() (int, error) {
	var selectedNumberOfTheMonth int
	fmt.Print("Seleziona il numero del mese in cui sei nato: ")
	_, err := fmt.Scanf("%d", &selectedNumberOfTheMonth)
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	return selectedNumberOfTheMonth, err
}

func getYear() (int, error) {
	var selectedYear int
	fmt.Print("In che hanno sei Nato?: ")
	_, err := fmt.Scanf("%d", &selectedYear)
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	return selectedYear, err
}

func getLastTwoDigitsOfTheYear(year int) (int, error) {
	yearAsString := strconv.Itoa(year)
	fmt.Println(yearAsString[:2])
	var lastTwoDigits int
	_, err := fmt.Sscan(yearAsString[2:4], &lastTwoDigits)
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	return lastTwoDigits, err
}

func main() {
	month, err := getMonths()
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	fmt.Println(month)

	day, err := getDayNumberOfTheMonth()
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	fmt.Println(day)

	year, err := getYear()
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	fmt.Println(year)
	yearLastTwoDigits, err := getLastTwoDigitsOfTheYear(year)
	if err != nil {
		log.Fatalf("errore %s", err)
	}
	fmt.Println("anno", yearLastTwoDigits)

	var sumDayMonth int = month + day
	if sumDayMonth >= 7 {
		sumDayMonth = sumDayMonth / 7
	}
	fmt.Println(sumDayMonth)

	if yearLastTwoDigits >= 7 {
		yearLastTwoDigits = yearLastTwoDigits / 7
	}
	fmt.Println("year last", yearLastTwoDigits)

	if yearLastTwoDigits >= 4 {
		yearLastTwoDigits = yearLastTwoDigits / 4
	}
	fmt.Println("quoziente", yearLastTwoDigits)

	// sommiamo

}
