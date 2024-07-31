package main

import (
	"fmt"
	"strings"
)

func main() {
	var letra string
	fmt.Println("Digite uma letra: ")
	fmt.Scan(&letra)

	letra = strings.ToLower(letra)
	if len(letra) != 1 || (letra[0] < 'a' || letra[0] > 'z') {
		fmt.Println("Caractere não aceito no programa.")
		return
	}

	switch letra {
	case "a", "e", "i", "o", "u":
		fmt.Println("A letra é uma vogal.")
	default:
		fmt.Println("A letra é uma consoante.")
	}
}
