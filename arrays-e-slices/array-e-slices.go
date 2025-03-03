package main

import (
	"fmt"
	"reflect"
)

// O tamanho de um array é fixo, ou seja, uma vez definido, não pode ser alterado.
//slice aponta para um array, ele parece que tem um poonteiro que aponta para um array ( slice é fatia em inglês)
//go é fortemente tipado slice e array nao interage entre si
// slice append pega um slice adiciona um novo item/valores e retorna um novo slice

// Resumo
// Array: Tamanho fixo, menos flexível, mas mais eficiente.
// Slice: Tamanho dinâmico, muito flexível, amplamente utilizado em Go.
// Slices são a forma preferida de trabalhar com coleções em Go devido à sua flexibilidade e facilidade de uso.
// Arrays são úteis em cenários onde o tamanho é conhecido e imutável.

func main() {
	fmt.Println("Array e Slices")
	fmt.Println("--------------------------")
	// slice
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(slice)
	fmt.Println("--------------------------")

	// slice append
	slice = append(slice, 110)
	fmt.Println("slice append->", slice)

	// array
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println("--------------------------")

	//slice referenciando um array criamos uma fatia de um array que já existe
	//imagine que isso é um ponteiro apontando para as posicoes
	//
	slice1 := array[1:3]

	fmt.Println(slice1)

	fmt.Println("--------------------------")

	// array
	var array1 [5]string
	array1[0] = "Posição 1"
	array1[1] = "Posição 2"
	array1[2] = "Posição 3"
	array1[3] = "Posição 4"
	array1[4] = "Posição 5"

	fmt.Println(array1)
	fmt.Println("--------------------------")

	fmt.Println("tipos->", reflect.TypeOf(slice))
	fmt.Println("tipos->", reflect.TypeOf(array))

	/*
		Funções úteis para Slices
		len(slice): Retorna o número de elementos no slice.
		cap(slice): Retorna a capacidade do slice (quantos elementos ele pode armazenar sem realocação).
		append(slice, elemento): Adiciona um ou mais elementos ao slice.
		Exemplo:
	*/
	fmt.Println("--- Funções úteis para Slices ---")
	sliceExemplo := []int{1, 2, 3}
	fmt.Println(len(sliceExemplo)) // Saída: 3
	fmt.Println(cap(sliceExemplo)) // Saída: 3

	sliceExemplo = append(sliceExemplo, 4, 5)
	fmt.Println(sliceExemplo)      // Saída: [1 2 3 4 5]
	fmt.Println(len(sliceExemplo)) // Saída: 5
	fmt.Println(cap(sliceExemplo)) // Saída: 6 (pode variar dependendo da realocação)
	fmt.Println("--------------------------")

	fmt.Println("--- Arrays Internos ---")
	// Arrays Internos
	// Internamente, um slice é representado por uma estrutura que contém um ponteiro para o array subjacente,
	// o tamanho do slice e a capacidade do array.
	// Quando você cria um slice a partir de um array, o slice aponta para o array existente.

	// Exemplo:
	arrayInterno := [5]int{1, 2, 3, 4, 5}
	sliceInterno := arrayInterno[1:3]
	// O sliceInterno aponta para o arrayInterno, compartilhando o mesmo array.
	// Modificando um elemento no sliceInterno também modifica o elemento no arrayInterno.
	sliceInterno[0] = 100

	fmt.Println(arrayInterno) // Saída: [1 100 3 4 5]

	/*
		A função make() em Go é usada para criar e inicializar certos tipos de dados internos, como slices, maps e channels.
		Ela é especialmente útil porque, além de alocar memória, ela também inicializa a estrutura de dados com valores padrão,
		garantindo que esteja pronta para uso.

		tipo: Tipo dos elementos do slice (por exemplo, int, string, etc.).
		tamanho: Número de elementos no slice.
		capacidade: Tamanho total do array subjacente (opcional). Se omitido, a capacidade é igual ao tamanho.
	*/

	fmt.Println("--- Arrays Internos make---")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println("Tamanho", len(slice3))    //tamanho
	fmt.Println("Capacidade", cap(slice3)) //capacidade
}
