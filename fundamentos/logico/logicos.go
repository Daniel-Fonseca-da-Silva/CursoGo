package main

import "fmt"

func compras(work1, work2 bool) (bool, bool, bool){
	comprarTV50 :=  work1 && work2
	comprarTV32 :=  work1 != work2 // ou exclusivo
	comprarSorvete := work1 || work2

	return comprarTV50, comprarTV32, comprarSorvete
}

func main(){
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
