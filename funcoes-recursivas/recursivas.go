package main

import "fmt"

//explique o que são funções recursivas
//funções recursivas são funções que chamam a si mesmas
//elas são úteis quando queremos executar uma função pequena e específica

func main() {
	fmt.Println(fibonacci(10))
}

// crie uma função recursiva que retorna um numero em uma determinada posicao de um fibonacci
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}
