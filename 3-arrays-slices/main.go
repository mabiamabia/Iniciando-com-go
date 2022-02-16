package main

import "fmt"

func main() {

	//arrays começam por posição 0, 1, ...
	const idade = 10
	var colors = [3]string{"Azul", "Vinho"}

	fmt.Println("O array tem", len(colors), "posições")
	fmt.Println("Os valores são", colors)
	fmt.Println("O primeiro valor é:", colors[0])

	colors[2] = "Vermelho"

	fmt.Println("O primeiro valor é:", colors[2])

	//slices
	var favoriteGames []string

	favoriteGames = append(favoriteGames, "The King of Fighters")
	favoriteGames = append(favoriteGames, "Bloodborn")
	favoriteGames = append(favoriteGames, "Street Fighter")
	favoriteGames = append(favoriteGames, "Super mario World")

	fmt.Println(favoriteGames)

}
