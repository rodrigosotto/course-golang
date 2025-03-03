package main

import "fmt"

/*
NO GO VOCE NAO PODE COMPARAR SOMAR ETC... VARIAEIS COM TIPOS DIFERENTES
*/
func main() {
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 2 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// NO GO VOCE NAO PODE COMPARAR, SOMAR ETC... VARIAEIS COM TIPOS DIFERENTES. go é fortemente tipado!
	var numero1 int16 = 10
	var numero2 int32 = 30

	//da pra especificar, transformar uma das variaveis no tipo ideal ( igual ) ai sim funcionaria,
	//mas creio nao ser uma boa pratica. Poderia fazer result := numero1 + int16(numero2) por exemplo
	result := numero1 + int16(numero2)
	fmt.Println(result)
	//FIM DOS ARITIMÉTICOS

	//OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "string1"
	variavel2 := "string2"

	fmt.Println(variavel1, variavel2)
	//FIM OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS ( sempre retorna um booleano)
	fmt.Println(1 > 0)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(2 == 2)
	fmt.Println(3 != 2)

	//FINAL OPERADORES RELACIONAIS

	//OPERADORES LOGICOS ( temos e ou ou negocao ou seja somente 3)

	fmt.Println("-------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println("-------------------------")
	//FINAL OEPRADORES LOGICO

	fmt.Println("-------------------------")

	numero := 10
	numero++

	numero += 15 // numero = numero + 15

	numero--
	numero -= 20 // numero = numero - 20

	numero = numero - 2
	numero /= 10
	numero %= 3
	numero *= 30 // numero = numero * 30

	//go nao aceita ternarios, para este tipo tem que ser if e else go so tem uma maneira de fazer as coisas neste caso ai tem que usar if else
	//texto := numero > 5 ? "Maior que 5" : "Menor que 5"

	var texto string
	if numero > 5 {
		fmt.Println("maior que 5")

	} else {
		fmt.Println("menor que 5")
	}
	fmt.Println(texto)
	fmt.Println("-------------------------")

}
