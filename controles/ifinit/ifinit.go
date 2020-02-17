package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// Parecida com a variavel dentro do for
	if i := numeroAleatorio(); i > 5 { // também suporta switch
		fmt.Println("Ganhou!!!")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu!")
		fmt.Println(i)
	}

}
