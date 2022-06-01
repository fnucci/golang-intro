package contas

import "fmt"

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente / Valor invalido", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor invalido", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) (string, float64, float64) {

	if valorDaTransferencia > 0 && valorDaTransferencia < c.Saldo {
		mensagemSaque, SaldoSaque := c.Sacar(valorDaTransferencia)
		fmt.Println(mensagemSaque, SaldoSaque)
		mensagemDeposito, SaldoDeposito := contaDestino.Depositar(valorDaTransferencia)
		fmt.Println(mensagemDeposito, SaldoDeposito)
		return "Transferencia realizada com sucesso", c.Saldo, contaDestino.Saldo
	} else {
		return "Saldo insuficiente / Valor invalido", c.Saldo, contaDestino.Saldo
	}
}
