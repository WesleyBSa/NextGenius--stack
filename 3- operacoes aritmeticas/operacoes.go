package main

import "fmt"

func main() {
	var a, b float64
	fmt.Println("Digite dois números: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Soma: %.2f\n", a+b)
	fmt.Printf("Subtração: %.2f\n", a-b)
	fmt.Printf("Multiplicação: %.2f\n", a*b)
	fmt.Printf("Divisão: %.2f\n", a/b)
}
