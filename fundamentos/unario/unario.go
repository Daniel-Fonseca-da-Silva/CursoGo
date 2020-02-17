package main

import "fmt"

func main(){
	x := 1
	y := 2

	// Apenas postfix
	x++ // x += 1 ou x = x + 1
	fmt.Println("x =", x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println("y =", y)

	// fmt.Println(x == y--)
}
