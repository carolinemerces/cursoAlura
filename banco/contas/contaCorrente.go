package contas

import (

	"github.com/carolinemerces/cursoAlura/banco/clientes"
)

type ContaCorrente struct {
	Titular cliente.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) (status string) {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Slado insuficiente!"
}

func (c *ContaCorrente) Depositar(deposito float64) (status string, saldo float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Depósito realizado com sucesso!", c.saldo
	}
	return "O valor do depósito menor que zero!", c.saldo
}

func (c *ContaCorrente) Transferir(transferencia float64, contaDestino *ContaCorrente) (status bool) {
	if transferencia < c.saldo && transferencia > 0 {
		c.saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() (saldo float64) {
	return c.saldo
}