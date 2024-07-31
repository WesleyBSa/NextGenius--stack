package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Println("Digite o valor de B: ")
	fmt.Scan(&b)

	a, b = b, a

	fmt.Printf("Após a troca, A = %d e B = %d\n", a, b)
}
