package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	exibeIntroducao()
	for {
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
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...!")
		default:
			fmt.Println("Opção invávila...!")
			sair(-1)
		}
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...!")
	site := "https://www.alura.com.br"
	resposta, error := http.Get(site)

	if resposta.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso")
	} else {
		fmt.Println("Site", site, "falhou no carregamento com o status code:", resposta.StatusCode, "com erro: ", error)
	}

	var sites [5]string
	sites[0] = "https://www.alura.com.br"
	sites[1] = "https://www.caelum.com.br"
	sites[2] = "https://www.youtube.com.br"
	sites[3] = "https://www.twitch.com.br"
	sites[4] = "https://www.gmail.com.br"
	fmt.Println(sites)

	exibeNomesComSlices()
}

func exibeNomesComSlices() {
	nomes := []string{"Felipe", "Adriana", "Edileia", "Ricardo"}
	fmt.Println(len(nomes))
	nomes = append(nomes, "Antonio", "Maria")
	fmt.Println(cap(nomes))
}
