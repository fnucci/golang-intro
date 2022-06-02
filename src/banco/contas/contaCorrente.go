package contas

import (
	"fmt"
	"golang-intro/src/banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "saldo insuficiente / Valor invalido", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor invalido", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) (string, float64, float64) {

	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		mensagemSaque, saldoSaque := c.Sacar(valorDaTransferencia)
		fmt.Println(mensagemSaque, saldoSaque)
		mensagemDeposito, saldoDeposito := contaDestino.Depositar(valorDaTransferencia)
		fmt.Println(mensagemDeposito, saldoDeposito)
		return "Transferencia realizada com sucesso", c.saldo, contaDestino.saldo
	} else {
		return "saldo insuficiente / Valor invalido", c.saldo, contaDestino.saldo
	}
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
