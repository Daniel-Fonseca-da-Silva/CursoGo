package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	// reflect.TypeOf -> informa o tipo
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Com sinal.... int8 int 16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa um mapeamento da tabela unicode(int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// Boolean

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá meu nome é Nikolai"
	fmt.Println(s1 + "!")
	// len -> informa tamanho da string
	fmt.Println("O tamanho da string é", len(s1))

	// String com multiplas linhas
	s2 := `Olá
	meu 
	nome
	é
	Nikolai
	`
	fmt.Println("O tamanho da string é", len(s2))

	// Char??? não existe
	// var x char = 'b' não existe
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Print(char)

}
