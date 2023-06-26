package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 6

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Printf("Digite um comando que seja válido!")
			//Indica que ocorreu algum erro
			os.Exit(-1)
		}

	}
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo...")
	// } else {
	// 	fmt.Printf("Digite um comando que seja válido!")
	// }
	// Para verificar o tipo da váriavel
	// fmt.Println("O tipo da variavel é", reflect.TypeOf(versao))
}

func exibeIntroducao() {
	nome := "Luiz"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	//	fmt.Println("O endereço da minha variável comando é", &comando)
	// fmt.Println("O comando escolhido foi,", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//slice
	// sites := []string{"https://www.alura.com.br",
	// 	"https://www.caelum.com.br", "https://www.google.com.br"}

	sites := leSitesDoArquivo()

	// fmt.Println(sites)
	for i := 0; i < monitoramentos; i++ {
		// range obtem o indice e o item correspondete do slice
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:",
			resp.StatusCode)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(arquivo)
	return sites
}
