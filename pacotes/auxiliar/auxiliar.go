package auxiliar

import (
	"fmt"
)

// escreve alguma coisa ( importante ser com letra Maiuscula para usar em outras partes do pacote ou seja ela é public)
// veja que nao foi necessario chamar/importar auxiliar.escrevendo2 porque escrevendo2 esta dentro da mesma pasta
// assim sendo ela pode ser importanda dentro do mesmo escopo algo assim.
func Escrever() {
	fmt.Println("Escrever do pacote auxiliar")
	escrevendo2()
}
