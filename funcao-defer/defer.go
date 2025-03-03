package main

import "fmt"

func main() {
	defer funcao1()
	funcao2()
	funcao3()
}

//exemplo de como usar o defer
//o defer é uma função que vai ser executada após a função main terminar
//ele é útil para fechar conexões, arquivos, etc...
//ele é executado na ordem inversa da chamada
//ele é executado mesmo que ocorra um erro

func funcao1() {
	fmt.Println("funcao1")
}

func funcao2() {
	fmt.Println("funcao2")
}

func funcao3() {
	fmt.Println("funcao3")
}
