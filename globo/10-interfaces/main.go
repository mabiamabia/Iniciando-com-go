package main

import "fmt"

//interface  == contrato
type transporter interface {
	calculateFreight(value float32) float32
}

type correios struct{}

func (c correios) calculateFreight(value float32) float32 {
	return value + 0.3
}

type fedex struct{}

func (f fedex) calculateFreight(value float32) float32 {
	return value + 0.5
}

func main() {
	fmt.Println("Hello world")
}
