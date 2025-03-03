package main

import "fmt"

func main() {
	total := soma(100, 200, 300)
	fmt.Println(total)

	resultado := print("Ola", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(resultado)
}

// explicando o que são funções variaticas
// funções variaticas são funções que podem receber um número variável de parametros
// 3 pontinhos significa que ela pode receber varios parametros de um mesmo tipo
// explicando o que é o ...int:
// o ...int significa que a função pode receber um número variável de inteiros
// exemplo: soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
// o total é a soma de todos os numeros passados como parametro
// função variatica com slice
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// crie um função que receba um texto e um conjunto de numero e retorne um string e print na tela
func print(texto string, numeros ...int) string {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return fmt.Sprintf("%s %d", texto, total)
}
