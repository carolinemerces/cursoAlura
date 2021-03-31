package contas

import cliente "github.com/carolinemerces/cursoAlura/banco/clientes"

type ContaPoupanca struct {
	Titular                  cliente.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(saque float64) (status string) {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Slado insuficiente!"
}

func (c *ContaPoupanca) Depositar(deposito float64) (status string, saldo float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Depósito realizado com sucesso!", c.saldo
	}
	return "O valor do depósito menor que zero!", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() (saldo float64) {
	return c.saldo
}

func (c *ContaPoupanca) Transferir(transferencia float64, contaDestino *ContaCorrente) (status bool) {
	if transferencia < c.saldo && transferencia > 0 {
		c.saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	}
	return false
}