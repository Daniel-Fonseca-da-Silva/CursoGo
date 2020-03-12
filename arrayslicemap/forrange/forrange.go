package main

import "fmt"

func main() {

	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	// (i) indice, (valor) numero
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// ( _ ) Ignora valor do indice
	for _, num := range numeros {
		fmt.Println(num)
	}
}
