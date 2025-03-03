package main

func main() {
	//exemplo de estrutura de controle IF ELSE
	idade := 12
	if idade >= 18 {
		println("Você é maior de idade")
	} else {
		println("Você é menor de idade")
	}

	//exemplo atribuindo um valor e verificando
	//quando voce cria uma variavel usando IF init voce esta limitando ela ao escopo do if
	//assim voce nao consegue acessar ela em outros lugares fora do escopo por exemplo

	if outraIdade := idade; outraIdade >= 18 {
		println("Você é maior de idade")
	} else {
		println("Você é menor de idade")
	}
	//exemplo de estrutura de controle else if
	if idade >= 18 {
		println("Você é maior de idade")
	} else if idade >= 16 {
		println("Você é adolescente")
	} else {
		println("Você é menor de idade")
	}

}
