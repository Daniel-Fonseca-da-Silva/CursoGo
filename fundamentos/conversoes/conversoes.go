package main

import (
	"fmt"
	c "strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	//
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado...
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste " + c.Itoa(123))

	// string para int
	// 1° valor(num), 2°error(ignorando erro '_')
	num, _ := c.Atoi("123")
	fmt.Println(num - 122)

	// converte bool

	b, _ := c.ParseBool("1")

	if b {
		fmt.Println("Verdadeiro")
	}

}
