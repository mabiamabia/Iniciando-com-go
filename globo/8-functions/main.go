package main

import "fmt"

func main() {

}

func sendMessages(msgs ...string) {
	for _, message := range msgs {

		fmt.Println("Enviando mensagem", message)
	}
}

func checkAge(name string, age uint) (string, bool) {
	if age >= 18 {

		return fmt.Sprintf(name, "Permitido"), true
	}
	//sprint retorna string que ele printaria
	return fmt.Sprint(name, "Não permitido"), false
}
