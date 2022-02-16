package main

import (
	"fmt"
)

func main() {

	var hobby = "musica"
	var name = "mabia"
	const idade = 30
	var firstName, lastName string = "ze claudio", "da silva junior"

	fmt.Println("Oi, ", name, "!", "Meu nome é", firstName, lastName, "e tenho", idade, "anos")
	fmt.Printf("Meu hobby é: %v \n", hobby)
}
