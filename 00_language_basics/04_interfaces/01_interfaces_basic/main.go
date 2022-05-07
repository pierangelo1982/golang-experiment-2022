package main

import "fmt"

type FullName interface {
	getFullName() string
}

type Person struct {
	FirstName, LastName string
}

func (p *Person) getFullName() string {
	//fmt.Printf("Full name is %s %s\n", p.FirstName, p.LastName)
	return p.FirstName + " " + p.LastName
}

func main() {
	p1 := &Person{
		FirstName: "Napoleone",
		LastName:  "Bonaparte",
	}

	p2 := &Person{
		FirstName: "Giuseppe",
		LastName:  "Garibaldi",
	}
	//p1.getFullName()
	fmt.Println(p1.getFullName())

	fmt.Println(p2.getFullName())
}
