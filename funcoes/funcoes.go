package main

import "fmt"

/*
ME GO AS FUNÇÕES PODE TER PARAMETROS E RETORNO ( PARAMETROS SAO DADOS QUE ADD NA FUNÇÃO E RETORNO É O QUE ELA DEVOLVE PRA GENTE)
EM GO ESPECIFICAMOS O RETORNO DEPOIS DO PARANTESES NO CASO A FUNCAO SOMAR RETORNA UM INT8
retono de uma função pode ser uma função por exemplo, posso igual uma variavel a uma funcao pois elas ( funcoes ) sao tipos.
em go voce pode  ter mais de 1 retorno e caso queira ignorar algum desses retorno basta user underline _
*/

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("dentro da funcao f ->", txt)
		return txt
	}

	resultadoDeF := f(" resultado de f--> texto como paramentro 1")
	fmt.Println(resultadoDeF)

	fmt.Println("---- Função calculos matemáticos ----")

	//caso queria mostrar somente 1 retorno voce adiciona _

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
