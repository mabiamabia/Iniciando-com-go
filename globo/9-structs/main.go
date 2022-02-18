package main

import "fmt"

type videoGame struct {
	name  string
	brand string
	year  uint
}

func (v videoGame) printInfo() {
	fmt.Println("Nome:", v.name, "Marca", v.brand, "Ano:", v.year)
}

func main() {
	playstation := videoGame{
		name:  "Play One",
		brand: "Sony",
		year:  2000,
	}

	fmt.Println(playstation)
	fmt.Println(playstation.name)
	playstation.printInfo()
}
