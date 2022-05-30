package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func main() {

	conta1 := ContaCorrente{
		titular: "Guilherme",
		agencia: 123,
		conta:   125433,
		saldo:   150.50,
	}

	conta2 := ContaCorrente{"Adriana", 324, 896745, 4.29}

	fmt.Println(conta1)

	fmt.Println(conta2)
}
