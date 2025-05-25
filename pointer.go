package main

import "fmt"

func main() {

	type Person struct {
		Name    string
		Age     uint8
		Address string
	}
	// Create pointer to struct
	p := &Person{
		Name: "Italo Rafael",
	}

	s := new(Person)
	s.Name = "Italo Rafael"

	fmt.Println(p)
	fmt.Println(s)
}
