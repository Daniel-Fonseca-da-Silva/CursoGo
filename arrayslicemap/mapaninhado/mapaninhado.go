package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1574.45,
			"Guga Pereira":   8456.23,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro de Alcantara": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
