package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Printf("Digite um comando que seja válido!")
			//Indica que ocorreu algum erro
			os.Exit(-1)
		}

	}
}

func exibeIntroducao() {
	nome := "Luiz"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
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
			fmt.Println("Testando site", i+1, ":", site)
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
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:",
			resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	//Abertura de arquivo
	arquivo, err := os.Open("sites.txt")

	// Tratamento de erro
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {

		//Leitura da linha do arquivp até a quebra de linha
		linha, err := leitor.ReadString('\n')

		//Remove a quebra de linha
		linha = strings.TrimSpace(linha)

		//Adicionando site no slice
		sites = append(sites, linha)

		//Forçando parada ao chegar no final do arquivo
		if err == io.EOF {
			break
		}
	}

	// Fechando arquivo
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	// Utiliza de flags para criar o arquivo de log, caso não tenha, escrever e adicionar linha a linha
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
