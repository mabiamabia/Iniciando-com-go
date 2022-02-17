package main

import "fmt"

func main() {
	points := 70

	if points >= 80 {
		fmt.Println(("Prova Top"))
	} else if points >= 60 {
		fmt.Println(("Bom"))
	} else if points >= 40 {
		fmt.Println(("Ruim"))
	}

	var dbzChars []string = []string{"goku"}

	dbzChars = append(dbzChars, "bulma")
	dbzChars = append(dbzChars, "trunks")
	dbzChars = append(dbzChars, "pan")
	dbzChars = append(dbzChars, "vegeta")
	dbzChars = append(dbzChars, "gohan")

	fmt.Println("Temos o goku?", contains(dbzChars, "goku"))

	/* ---------------------------- */

	fruta := "kiwi"

	switch fruta {
	case "uva":
		var a = 1
		a++
		fmt.Println(a)
	case "kiwi":
		fmt.Println(gosto nem)
	default:
		fmt.Println("_")

	}

}

func contains(dbzChars []string, s string) {
	panic("unimplemented")
}
