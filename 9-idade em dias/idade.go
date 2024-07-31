package main

import "fmt"

func main() {
	var anos, meses, dias int
	fmt.Println("Digite sua idade em anos, meses e dias: ")
	fmt.Scan(&anos, &meses, &dias)

	totalDias := (anos * 365) + (meses * 30) + dias
	fmt.Printf("Sua idade em dias é: %d\n", totalDias)
}
