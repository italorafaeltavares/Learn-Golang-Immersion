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

	// slices
	s := make([]string, 0)
	s = append(s, "Italo")
	s = append(s, "Rafael")

	s2 := make([]int8, 2)
	s2[0] = 1
	s2[1] = 2
	s2 = append(s2, 3)

	//array
	var a [2]bool
	a[0] = true
	a[1] = false

	fmt.Println(name, age, document)
	fmt.Println(p)
	fmt.Println(s, s2, len(s2))
	fmt.Println(a, len(a))
}
