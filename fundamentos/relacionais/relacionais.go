package main

import "fmt"
import "time"

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 3)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println(">", 3 <= 2)
	fmt.Println(">", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	// fmt.Println("Datas:", d1.Equals(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Pedro"}
	p2 := Pessoa{"Pedr"}

	fmt.Println("Pessoas:", p1 == p2)
}
