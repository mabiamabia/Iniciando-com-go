package main

import "fmt"

func main() {
	languageLevel := make(map[string]int)

	languageLevel["Java"] = 10
	languageLevel["PHP"] = 5
	languageLevel["Javascript"] = 9
	languageLevel["GoLang"] = 7
	languageLevel["Dart"] = 3

	fmt.Println(languageLevel)

	languageLevel["Dart"] = 8

	fmt.Println(languageLevel)

	delete(languageLevel, "Dart")

	fmt.Println(languageLevel)

	fmt.Println("quantidade", len(languageLevel))

	/* ---------------------------------------------------------------- */

	consoleAndGames := make(map[string][]string)

	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Donkey Kong")
	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Super Mario World")
	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"], "Star Fox")

	consoleAndGames["Mega Drive"] = append(consoleAndGames["Mega Drive"], "Sonic")
	consoleAndGames["Mega Drive"] = append(consoleAndGames["Mega Drive"], "Streets of Rage")
	consoleAndGames["Mega Drive"] = append(consoleAndGames["Mega Drive"], "Shinobi")

	fmt.Println(consoleAndGames)
	fmt.Println(consoleAndGames["Super Nintendo"])
	fmt.Println(consoleAndGames["Super Nintendo"][1])

	consoleAndGames["Super Nintendo"] = append(consoleAndGames["Super Nintendo"][:1], consoleAndGames["Super Nintendo"][2:3]...)

}
