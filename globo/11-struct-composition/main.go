package main

import "fmt"

type cliente struct {
	nome string
	email string
	endereco
}

type endereco struct{
	rua string
	numero uint

}

func main() {

	jose := cliente {
		nome: "Optimus Prime",
		email: "optimusprime@gmail.com"
		endereco: endereco {
			rua: "Dark side of the moon"
			numero: "43"
		}
	},
}
