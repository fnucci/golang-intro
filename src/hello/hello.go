package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	// if escolha == 1 {
	// 	fmt.Println("Monitorando...!")
	// } else if escolha == 2 {
	// 	fmt.Println("Exibindo logs...!")
	// } else if escolha == 0 {
	// 	fmt.Println("Saindo do programa...!")
	// } else {
	// 	fmt.Println("Opção invávila...!")
	// }

	opcaoEscolhida := lerEntrada()

	switch opcaoEscolhida {
	case 0:
		fmt.Println("Saindo do programa...!")
		sair(0)
	case 1:
		fmt.Println("Monitorando...!")
	case 2:
		fmt.Println("Exibindo logs...!")
	default:
		fmt.Println("Opção invávila...!")
		sair(-1)
	}

}

func exibeIntroducao() {
	nome := "Felipe"
	idade := 32
	versao := 1.2

	fmt.Println("Hello World GO")
	fmt.Println("Olá sr. ", nome, " sua idade é ", idade)
	fmt.Println("Programa na versão ", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))

}

func exibeMenu() {
	fmt.Println("1- Monitorar sistemas")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func lerEntrada() int {

	var escolha int
	fmt.Scan(&escolha)

	fmt.Println("voce escolheu a opcao", escolha)
	fmt.Println("o endereco da variavel escolha é", &escolha)

	return escolha
}

func sair(codigo int) {
	os.Exit(codigo)
}
