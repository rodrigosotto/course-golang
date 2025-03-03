package main

// exemplo de retorno nomeado
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// exemplo sem retorno nomeado
func calculosMatematicos2(n1, n2 int) (int, int) {
	soma2 := n1 + n2
	subtracao2 := n1 - n2
	return soma2, subtracao2
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	println(soma, subtracao)

	soma2, subtracao2 := calculosMatematicos2(10, 20)
	println(soma2, subtracao2)
}
