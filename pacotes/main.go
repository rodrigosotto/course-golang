package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	const lca bool = true
	const name string = "Lucas"
	if lca {
		auxiliar.Escrever()
		checkmail.ValidateFormat("jrsottodev@gmail.com")
	} else {
		fmt.Println("Nao é", name)
	}
}
