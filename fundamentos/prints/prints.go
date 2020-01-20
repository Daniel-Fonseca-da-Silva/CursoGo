package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)

	// Sprint -> retorna em formato de String
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x, "!!!")
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Epa!"

	// \n -> break row
	// %d -> valor numerico
	// %f -> valor flutuante
	// %.1f -> valor flutuante
	// %t -> valor booleano
	// %d -> valor Character

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
