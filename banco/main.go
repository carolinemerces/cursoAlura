package main

import (
	"fmt"

	"github.com/carolinemerces/cursoAlura/banco/contas"
	"github.com/carolinemerces/cursoAlura/banco/clientes"
)

func main() {

	cliente1 := cliente.Titular{"Caroline", "123.456.789-90", "Desenvolvedora Go"}
	conta1 := contas.ContaCorrente{cliente1, 123, 123456, 300}
	
	fmt.Println(conta1)

	conta2 := new(contas.ContaCorrente)
	conta2.Titular.Nome = "Camila"
	conta2.Titular.Profissao = "Desenvolvedora Go"
	//conta2.Saldo = 230.99
	fmt.Println(*conta2)

	// saque := 20.00
	// conta2.Saldo = conta2.Saldo - saque
	// fmt.Println(conta2.Saldo)

	fmt.Printf("Seu Saldo: %2.f\n", conta2.Saldo)
	fmt.Println(conta2.Sacar(50.00))

	fmt.Printf("Seu Saldo: %2.f\n", conta2.Saldo)
	fmt.Println(conta2.Depositar(60.00))

	status := conta1.Transferir(45.87, conta2)
	fmt.Println(status)
	fmt.Println(conta1.Saldo, conta2.Saldo)
}
