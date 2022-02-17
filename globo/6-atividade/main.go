package main

import "fmt"

func main() {
	heroMarvel := make(map[string]int)

	heroMarvel["Homem Aranha"] = 7
	heroMarvel["Super Homem"] = 8
	heroMarvel["Batman"] = 6
	heroMarvel["Super Shock"] = 5

	fmt.Println(heroMarvel)

	for _, e := range heroMarvel {
		if e >= 7 {
			fmt.Println("Muito poderoso uau")
		} else if 7 <= 5 {
			fmt.Println("Poderosinho")
		} else {
			fmt.Println("Muito fraquinho, tadinho")
		}

	}
}
