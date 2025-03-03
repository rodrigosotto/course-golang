package main

import (
	"fmt"
	"time"
)

func main() {
	//exemplo de loops com o go
	i := 0
	for i < 3 {
		i++
		println("encrementando I: ", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 2; j++ {
		fmt.Println("encrementando J: ", j)
		time.Sleep(time.Second)
	}

	//exemplo de for com a clausula range
	nomes := [3]string{"Jose", "Manuel", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(nome)
		fmt.Println(indice)
	}

	for indice, letra := range "palavra" {
		fmt.Println(indice, string(letra))
	}

	//exemplo de interação com map
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
