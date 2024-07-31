package main

import "fmt"

func main() {
	var golsTime1, golsTime2 int
	fmt.Println("Digite os gols do time 1: ")
	fmt.Scan(&golsTime1)
	fmt.Println("Digite os gols do time 2: ")
	fmt.Scan(&golsTime2)

	if golsTime1 > golsTime2 {
		fmt.Println("Time 1 venceu!")
	} else if golsTime2 > golsTime1 {
		fmt.Println("Time 2 venceu!")
	} else {
		fmt.Println("Empate!")
	}
}
