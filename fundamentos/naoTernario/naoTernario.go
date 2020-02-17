package main

import "fmt"

// Não tem operador ternário em go ~.~

func obterResultado(nota float64) string {

	// // Notação de um operador ternário
	// return nota >= 6 ? "Aprovado" : "Reprovado"

	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}

func main() {
	fmt.Println(obterResultado(6.2))
}
