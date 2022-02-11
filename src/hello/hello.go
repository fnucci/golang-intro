package main

import (
	"fmt"
	"reflect"
)

func main() {

	nome := "Felipe"
	idade := 32
	versao := 1.2

	fmt.Println("Hello World GO")
	fmt.Println("Olá sr. ", nome, " sua idade é ", idade)
	fmt.Println("Programa na versão ", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))

	fmt.Println("1- Monitorar sistemas")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")

	var escolha int
	fmt.Scan(&escolha)

	fmt.Println("voce escolheu a opcao", escolha)
	fmt.Println("o endereco da variavel escolha é", &escolha)

	// if escolha == 1 {
	// 	fmt.Println("Monitorando...!")
	// } else if escolha == 2 {
	// 	fmt.Println("Exibindo logs...!")
	// } else if escolha == 0 {
	// 	fmt.Println("Saindo do programa...!")
	// } else {
	// 	fmt.Println("Opção invávila...!")
	// }

	switch escolha {
	case 0:
		fmt.Println("Saindo do programa...!")
	case 1:
		fmt.Println("Monitorando...!")
	case 2:
		fmt.Println("Exibindo logs...!")
	default:
		fmt.Println("Opção invávila...!")
	}

}
