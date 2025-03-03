package main

//NO GO SWITCH NÃO TEM BREAKS, POIS ELA SAI AUTOMATICAMENTE

func diaDaSemama(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}
func main() {
	//exemplo de estrutura de controle SWITCH
	idade := 17
	switch {
	case idade >= 18:
		println("Você é maior de idade")
	case idade >= 16:
		println("Você é menor de idade mas pode votar")
	}

	dia := diaDaSemama(1)
	println(dia)

}
