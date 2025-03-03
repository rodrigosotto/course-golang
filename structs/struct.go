package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	cep    string
	numero int16
}

func main() {
	fmt.Println("exemplos de structs")
	fmt.Println("--------------")

	var user usuario
	user.idade = 21
	user.nome = "Jefferson"

	fmt.Println(user)
	fmt.Println("--------------")

	enderecoExemplo := endereco{rua: "Flor de Lis", cep: "85855-000", numero: 100}

	fmt.Println("user2--------------")

	user2 := usuario{"Julius", 22, enderecoExemplo}

	fmt.Println(user2)
}
