package main

import (
	"fmt"

	"github.com/carolinemerces/cursoAlura/banco/contas"
	//"github.com/carolinemerces/cursoAlura/banco/clientes"
)

func PagarBoleto(conta verificarContaBoleto, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarContaBoleto interface {
	Sacar(valor float64) string
}

func main() {

	contaExemplo := contas.ContaPoupanca{}
	conta1 := contas.ContaCorrente{}

	contaExemplo.Depositar(200)
	conta1.Depositar(300)

	PagarBoleto(&contaExemplo, 70)
	PagarBoleto(&conta1, 78.99)

	fmt.Println(contaExemplo.ObterSaldo(), conta1.ObterSaldo())
}
