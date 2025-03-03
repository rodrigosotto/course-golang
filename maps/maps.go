package main

import "fmt"

/*
Um map é uma coleção de pares chave-valor. Quando você cria um map usando make(),
ele inicializa a estrutura interna do map, permitindo que você comece a adicionar pares chave-valor.
O map é inicializado e pronto para uso.
Se você tentar acessar uma chave que não existe, o valor zero do tipo do valor será retornado (por exemplo, 0 para int).
*/

func main() {
	//exemplo de map
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	//para acessar uma chave
	fmt.Println(usuario["nome"])

	//para adicionar uma chave
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	//exemplo para deletar uma chave do usuario2
	//delete é uma funcao nativa do GO
	delete(usuario2, "nome")

	fmt.Println(usuario2)

	//exemplo para adicionar uma chave do usuario2
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)

	//exemplo para percorrer um map
	for chave, valor := range usuario2 {
		fmt.Println("Percorrendo->", chave, valor)
	}

}
