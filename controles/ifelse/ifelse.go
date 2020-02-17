package main

import "fmt"

func imprimirResultador(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultador(7.3)
	imprimirResultador(5.1)
}
