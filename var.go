package main

import "fmt"

// var idade uint8 = 32
// var documento string = "000.000.000-00"

var (
	age      uint8  = 32
	document string = "000.000.000-00"
)

// struct
type Person struct {
	Name     string
	Age      uint8
	Document string
	Address  string
	CEP      string
}

func main() {
	// var nome string = "Italo Rafael"
	name := "Italo Rafael"

	p := Person{
		Name:     "Italo Rafael",
		Age:      37,
		Document: "000.000.000-00",
		Address:  "Rua 01",
	}

	fmt.Println(name, age, document)
	fmt.Println(p)
}
