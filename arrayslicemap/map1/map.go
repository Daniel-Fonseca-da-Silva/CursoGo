package main

import "fmt"

func main() {

	// var aprovados map[int]string
	// Mapas devem ser incializados!!!!

	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[23523567] = "Pedro"
	aprovados[98663456] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s CPF: %d\n", nome, cpf)
	}

	fmt.Println(aprovados)

}
