package main

import "fmt"

func main(){
	i := 1

	// Go não tem aritmética de ponteiro!
	var p *int = nil // nil = (null)nulo!, criando ponteiro

	p = &i // pegando o endereço da variável
	*p++ // desreferenciando (pegando valor)
	i++

	// p++

	fmt.Println(p, *p, i, &i)
}
