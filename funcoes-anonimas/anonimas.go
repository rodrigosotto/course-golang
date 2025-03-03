package main

import "fmt"

// funções anonimas são funções que não tem nome
// elas são definidas e imediatamente executadas
// elas são úteis quando queremos executar uma função pequena e específica
// elas são úteis quando queremos executar uma função pequena e específica
// elas são úteis quando queremos executar uma função pequena e específica

func main() {
	func() {
		fmt.Println("Ola, mundo!")
	}() //exemplo de como deixar a funcao anonima em uma linha

	func(texto string) {
		fmt.Println(texto)
	}("passando o parametros") // aqui passo o parametro que ela vai retornar

	// aqui eu defino uma funcao anonima e retorno uma string
	// e passo o parametro que ela vai retornar
	// e atribuo o retorno a uma variavel
	// e imprimo a variavel
	// veja que é possivel atribuir o retorno a uma variavel
	// e imprimir a variavel

	//sprintf é uma funcao do fmt que pode ser usada para formatar uma string, concatenar strings e etc...
	//para cada tipo de dado eu preciso passar %s, %d, %f, %t, %v, etc...
	//%s é para string
	//%d é para int
	//%f é para float
	//%t é para bool
	//%v é para qualquer tipo de dado
	//%s é para string
	retorno := func(texto string) string {
		return fmt.Sprintf("retorno: %s", texto)
	}("passando o parametros 2") // aqui passo o parametro que ela vai retornar

	fmt.Println(retorno)
}
