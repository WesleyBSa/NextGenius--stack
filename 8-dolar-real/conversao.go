package main

import "fmt"

func main() {
	var valorDolar, cotacaoDolar, valorReal float64
	fmt.Println("Digite o valor em dólares (US$): ")
	fmt.Scan(&valorDolar)
	fmt.Println("Digite a cotação do dólar: ")
	fmt.Scan(&cotacaoDolar)

	valorReal = valorDolar * cotacaoDolar
	fmt.Printf("O valor em reais (R$) é: %.2f\n", valorReal)
}
