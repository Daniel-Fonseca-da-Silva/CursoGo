package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não existe while em Go
		fmt.Println(i)
		i++
	}

	// For tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// Mistureba |~.~|
	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par!")
		} else {
			fmt.Print("Impar!")
		}
	}
	// Laço infinito |-_-|
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
