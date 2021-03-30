package main

import (
	"fmt"

	"github.com/carolinemerces/cursoAlura/banco/contas"
	//"github.com/carolinemerces/cursoAlura/banco/clientes"
)

func main() {

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(200)
	fmt.Println(contaExemplo.ObterSaldo())
}
