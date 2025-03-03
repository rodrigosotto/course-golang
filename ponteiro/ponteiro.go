package main

import "fmt"

// Podemos pensar que ponteiros é uma variavel que vai salvar dentro dela um endereço de memoria e não um valor
// Quando você atibui um valor de uma variavel por padrao é uma copia do valor para outra variavel
// Ponteiro é uma referencia de memoria
func main() {
	//exemplo de ponteiros
	fmt.Println("--- Exemplo de Ponteiros ---")

	//variavel a
	a := 10
	//variavel ponteiro b que vai guardar o endereço de memoria da variavel a
	var b *int = &a
	//imprimindo o endereço de memoria da variavel a
	fmt.Println(&a)
	//imprimindo o valor da variavel a
	fmt.Println(a)
	//imprimindo o endereço de memoria da variavel b
	fmt.Println(b)
	//imprimindo o valor da variavel b
	fmt.Println(*b)

}
