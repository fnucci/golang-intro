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

	conta2 := ContaCorrente{"Adriana", 324, 896745, 114.29}

	var conta3 *ContaCorrente
	conta3 = new(ContaCorrente)

	conta3.titular = "Mocreia"
	conta3.agencia = 543
	conta3.conta = 989870
	conta3.saldo = 300.67

	cloneConta1 := ContaCorrente{
		titular: "Guilherme",
		agencia: 123,
		conta:   125433,
		saldo:   150.50,
	}

	cloneConta2 := ContaCorrente{"Adriana", 324, 896745, 4.29}

	var cloneConta3 *ContaCorrente
	cloneConta3 = new(ContaCorrente)

	cloneConta3.titular = "Mocreia"
	cloneConta3.agencia = 543
	cloneConta3.conta = 989870
	cloneConta3.saldo = 300.67

	fmt.Println("Conteudos: ", conta1, conta2, *conta3)

	fmt.Println("Enderecos: ", &conta1, &conta2, &conta3)

	fmt.Println("Compare conta 1: ", conta1 == cloneConta1, "Enderecos: ", &conta1 == &cloneConta1)
	fmt.Println("Compare conta 2: ", conta2 == cloneConta2, "Enderecos: ", &conta2 == &cloneConta2)
	fmt.Println("Compare conta 3: ", *conta3 == *cloneConta3, "Enderecos: ", &conta3 == &cloneConta3)

	fmt.Println(conta1.Sacar(50))
	fmt.Println("Saldo atualizado: ", conta1.saldo)
	fmt.Println(conta2.Sacar(75))
	fmt.Println("Saldo atualizado: ", conta2.saldo)
	fmt.Println(conta3.Sacar(150))
	fmt.Println("Saldo atualizado: ", conta3.saldo)

}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente / Valor invalido"
	}
}
