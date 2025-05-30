package main

import "fmt"

func main() {
	fmt.Println("Conditional in Go")

	animal := "cachorro"

	if animal == "gato" {
		println("mial")
	} else if animal == "cachorro" {
		println("au au")
	} else if animal == "passaro" {
		println("piu piu")
	} else {
		println("desconhecido")
	}

	switch animal {
	case "gato":
		println("mial")
	case "cachorro":
		println("au")
	case "passaro":
		println("piu")
	default:
		println("desconhecido")
	}
	idadeCrianca := 4
	alturaCrianca := 150

	if idadeCrianca >= 6 {
		println("Pode brinacar")
	} else if idadeCrianca >= 4 && alturaCrianca > 150 {
		println("Pode brincar acompanhdo!")
	} else {
		println("Não pode brincar")
	}

	switch {
	case idadeCrianca >= 6:
		println("Pode brinacar")
	case idadeCrianca >= 4 && alturaCrianca > 150:
		println("Pode brincar acompanhdo!")
	default:
		println("Não pode brincar")
	}
}
