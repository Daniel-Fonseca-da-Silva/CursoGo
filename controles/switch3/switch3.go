package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case bool:
		return "Boleano"
	case func():
		return "Funcao"
	default:
		return "Tipo desconhecido"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo(true))
	fmt.Println(tipo("Hallo!"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
