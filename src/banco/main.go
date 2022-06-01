package main

import (
	"fmt"

	"golang-intro/src/banco/contas"
)

func main() {

	conta1 := contas.ContaCorrente{
		Titular: "Guilherme",
		Agencia: 123,
		Conta:   125433,
		Saldo:   150.50,
	}

	conta2 := contas.ContaCorrente{"Adriana", 324, 896745, 114.29}

	var conta3 *contas.ContaCorrente
	conta3 = new(contas.ContaCorrente)

	conta3.Titular = "Mocreia"
	conta3.Agencia = 543
	conta3.Conta = 989870
	conta3.Saldo = 300.67

	cloneConta1 := contas.ContaCorrente{
		Titular: "Guilherme",
		Agencia: 123,
		Conta:   125433,
		Saldo:   150.50,
	}

	cloneConta2 := contas.ContaCorrente{"Adriana", 324, 896745, 4.29}

	var cloneConta3 *contas.ContaCorrente
	cloneConta3 = new(contas.ContaCorrente)

	cloneConta3.Titular = "Mocreia"
	cloneConta3.Agencia = 543
	cloneConta3.Conta = 989870
	cloneConta3.Saldo = 300.67

	var novoSaldo float64 = 0.0

	var mensagemRetorno string = ""

	var saldoOrigem float64 = 0.0
	var saldoDestino float64 = 0.0

	fmt.Println("Conteudos: ", conta1, conta2, *conta3)

	fmt.Println("Enderecos: ", &conta1, &conta2, &conta3)

	fmt.Println("Compare conta 1: ", conta1 == cloneConta1, "Enderecos: ", &conta1 == &cloneConta1)
	fmt.Println("Compare conta 2: ", conta2 == cloneConta2, "Enderecos: ", &conta2 == &cloneConta2)
	fmt.Println("Compare conta 3: ", *conta3 == *cloneConta3, "Enderecos: ", &conta3 == &cloneConta3)

	mensagemRetorno, novoSaldo = conta1.Sacar(50)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldo atualizado: ", novoSaldo)

	mensagemRetorno, novoSaldo = conta2.Sacar(75)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldo atualizado: ", novoSaldo)

	mensagemRetorno, novoSaldo = conta3.Sacar(150)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldo atualizado: ", novoSaldo)

	mensagemRetorno, novoSaldo = conta1.Depositar(500)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldo atualizado: ", novoSaldo)

	mensagemRetorno, novoSaldo = conta2.Depositar(600)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldo atualizado: ", novoSaldo)

	mensagemRetorno, novoSaldo = conta3.Depositar(750)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldo atualizado: ", novoSaldo)

	mensagemRetorno, saldoOrigem, saldoDestino = conta1.Transferir(40, conta2)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldos atualizados - Origem: ", saldoOrigem, " - Destino: ", saldoDestino)

	mensagemRetorno, saldoOrigem, saldoDestino = conta2.Transferir(100, *conta3)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldos atualizados - Origem: ", saldoOrigem, " - Destino: ", saldoDestino)

	mensagemRetorno, saldoOrigem, saldoDestino = conta3.Transferir(90, conta1)
	fmt.Println(mensagemRetorno)
	fmt.Println("Saldos atualizados - Origem: ", saldoOrigem, " - Destino: ", saldoDestino)
}
