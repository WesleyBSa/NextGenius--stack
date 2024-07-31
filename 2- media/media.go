package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Println("Digite três números: ")
	fmt.Scan(&n1, &n2, &n3)
	media := (n1 + n2 + n3) / 3
	fmt.Printf("A média é: %.2f\n", media)
}
