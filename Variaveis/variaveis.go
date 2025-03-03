package main

import (
	"fmt"
	"time"
)

//ANOTAÇOES
/*
GO É FORTEMENTE TIPADO, PODEMOS DECLARAR VARIAVEIS COM TIPOS EXPLICITO OU INPLICITO
GO NÃO DEIXA VOCE CRIAR UMA VARIAVEL E NAO USAR, VAI APARECER UM ERRO, MESMA COISA PARA UM IMPORT
GO EM TODO TIPO DE DADOS TEM UM VALOR INICIAL QUE SERIA 0
GO TEM UM TIPO ESPECIFICO PARA ERROS, ERROS TEM TIPO
EM GO, O TIPO DE DADO PARA REPRESENTAR DATAS E HORAS É time.Time, QUE FAZ PARTE DO PACOTE TIME NATIVO DO GO

DATE:
Em Go, para lidar com datas no formato "dd/mm/aaaa", você precisa usar o pacote time e seu método time.Parse(),
lembrando que Go usa uma referência fixa de formatação baseada na data "02/01/2006".
*/

func main() {
	// tipo explicito adicionei o tipo ( string )
	var variavel1 string = "variavel1"

	//inferindo o tipo de acordo com o valor dele
	variavel2 := "Variavel2"

	//declarar varias variaveis
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	//outra maneira de declar
	variavel5, variavel6 := "variavel5", "variavel6"

	//declar constante que o valor nao pode mudar
	const variavel7 = "constante variavel7"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(variavel7)

	// Obtendo a data e hora atual
	agora := time.Now()
	fmt.Println("Data e hora atual:", agora)

	// Criando uma data específica
	dataEspecifica := time.Date(2024, time.February, 1, 14, 30, 0, 0, time.UTC)
	fmt.Println("Data específica:", dataEspecifica)

	// Formatando a data
	formato := "02/01/2006 15:04:05"
	fmt.Println("Data formatada:", agora.Format(formato))

	dataStr := "01/02/2024 14:30:00"
	layout := "02/01/2006 15:04:05"

	dataConvertida, err := time.Parse(layout, dataStr)
	if err != nil {
		fmt.Println("Erro ao converter data:", err)
	} else {
		fmt.Println("Data convertida:", dataConvertida)
	}

	agora2 := time.Now()
	dataFormatada := agora2.Format("02/01/2006")
	fmt.Println("Data formatada:", dataFormatada)

	//ADICIONAR DIAS EM UMA DATA ESPECIFICA
	novaData := dataConvertida.AddDate(0, 0, 5) // Adiciona 5 dias
	fmt.Println("Nova data:", novaData.Format("02/01/2006"))

}
