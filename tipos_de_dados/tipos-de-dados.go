package main

import (
	"errors"
	"fmt"
)

/*
ANOTAÇÕES
GO TEM UM PACOTE NATIVO PARA ERROS
ERRO SEMPRE RETORNA <NIL> SEJA ELE PARA PONTEIROS POR EXEMPLO
NO GO ERRO NAO PODE SER ATRIBUIDO STRING POR EXEMPLO PORQUE ELE É DO TIPO ERRO E O TIPO ERRO TEM SEU TIPO.
NO GO PARA CRIAR UM ERRO OU SEJA ATRIBUIR UM VALOR PARA ERROS VOCE TEM QUE USAR O PACOTE ERROS AI SIM DA PRA ATRIBUIR

Go não tem um tipo object como em JavaScript, mas usa structs para definir tipos personalizados com múltiplos campos:
*/

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	var boleano1 bool
	fmt.Println(boleano1)

	var erro error = errors.New("ERRO INTERNO")
	fmt.Println(erro)

	var numero0 int8 = 123
	fmt.Println(numero0)

	var numero1 int16 = 12345
	fmt.Println(numero1)

	var numero2 int32 = 123123123
	fmt.Println(numero2)

	var numero3 int64 = 1239977693565229
	fmt.Println(numero3)

	//tipos int tem alias
	var numero4 rune = 123
	fmt.Println(numero4)

	var numero5 byte = 112
	fmt.Println(numero5)

	var texto string = "um texto qualquer"
	fmt.Println(texto)

	var numero6 float32 = 1.34823
	fmt.Println(numero6)

	var numero7 float64 = 164333363.98
	fmt.Println(numero7)

	//TIPOS ARRAY
	var numero8 [8]int //array de inteiros com tamanho fixo de 8 posições
	numero8[0] = 20
	numero8[1] = 10
	numero8[2] = 20
	numero8[3] = 10
	numero8[4] = 20
	numero8[5] = 10
	numero8[6] = 20
	numero8[7] = 10
	fmt.Println(numero8)

	//se precisar uma coleção dinamica o slice é bastante usado
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//NO GO STRUCT É EQUIVALENTE A UM OBJETO
	var pessoa Pessoa
	pessoa.Nome = "Jefferson"
	pessoa.Idade = 41

	fmt.Println(pessoa)

	//Se precisar de algo mais dinâmico como um mapa de chave/valor, pode usar um map:
	//Se precisar de flexibilidade, o map pode simular objetos dinâmicos,
	//enquanto structs são mais performáticas e seguras.

	obj := map[string]interface{}{
		"nome":  "Rodrigo",
		"idade": 41,
	}
	search := "Rodrigo"
	//percorrendo com FOR
	for key, value := range obj {
		if value == search {
			fmt.Printf("Encontrado! A chave é '%s' e o valor é '%s'\n", key, value)
			break
		}
	}
	fmt.Println(obj)
	//Se souber a chave podemos acessar diretamente.
	if nome, existe := obj["nome"]; existe {
		if nome == "Rodrigo" {
			fmt.Println("Rodrigo encontrado!")
		}
	}

}
