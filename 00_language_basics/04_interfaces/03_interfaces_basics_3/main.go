package main

import (
	"fmt"
	"strings"
)

type lowercase interface {
	setLowercase() string
	setCapitaliseName() string
}

type Person struct {
	Name string
}

func (p *Person) setLowercaseName() string {
	return strings.ToLower(p.Name)
}

func (p *Person) setCapitaliseName() string {
	return strings.Title(p.Name)
}

func main() {
	person := Person{Name: "GEORGE WASHINGTON"}
	fmt.Println(person.setLowercaseName())
	person2 := Person{Name: person.setLowercaseName()}
	fmt.Println(person2.setCapitaliseName())
}
