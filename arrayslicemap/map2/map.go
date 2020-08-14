package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"José João":    11325.45,
		"Fernanda Nej": 15456.78,
		"Pedro Junior": 1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Não existe")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
