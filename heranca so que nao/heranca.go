package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	var p1 pessoa
	p1.nome = "Joao"
	p1.idade = 30
	p1.sobrenome = "Da Silva"
	p1.altura = 200

	var p2 pessoa
	p2.nome = "Chico"
	p2.sobrenome = "Da Cunha"
	p2.idade = 41
	p2.altura = 190

	var estudante1 estudante
	estudante1.pessoa = p1
	estudante1.curso = "TADS"
	estudante1.faculdade = "IFPR"

	fmt.Println("--------------")

	fmt.Println("ESTUDANTE 1-->", estudante1)

	fmt.Println("--------------")

	//estudante2 := estudante{p1, "engenharia", "PUC"}

	var estudante2 estudante
	estudante2.pessoa = p2
	estudante2.curso = "Analise de dados"
	estudante2.faculdade = "PUC MINAS"

	fmt.Println("ESTUDANTE2-->", estudante2)

	fmt.Println("--------------")

	p3 := pessoa{"Marcolino", "Rocha", 32, 170}

	estudante3 := estudante{p3, "Analista de dados", "Unioeste"}

	fmt.Println("ESTUDANTE-->", estudante3)

	fmt.Println("--------------")
}
