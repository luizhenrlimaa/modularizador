package main

import (
	"fmt"
)

func main() {
	nome := "Luiz"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do programa")

	// fmt.Println("O tipo da variavel é", reflect.TypeOf(versao))
}
