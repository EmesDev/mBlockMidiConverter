package main

import (
	"bufio"
	"fmt"
	"mBlockConverter/services"
	"os"
)

// Função principal
func main() {
	// Definindo cores ANSI
	reset := "\033[0m"
	blue := "\033[34m"
	cyan := "\033[36m"
	yellow := "\033[33m"
	green := "\033[32m"
	red := "\033[31m"

	// Cabeçalho
	fmt.Println(yellow + "-------------------------------" + reset)
	fmt.Println(blue + "Linkedin: " + cyan + "https://www.linkedin.com/in/emesdev/" + reset)
	fmt.Println(blue + "Software feito por: " + cyan + "https://github.com/EmesDev" + reset)
	fmt.Println(blue + "Licença MIT" + reset)
	fmt.Println(yellow + "-------------------------------" + reset)

	// Obtendo e listando as portas seriais
	services.GetPorts()
	fmt.Println(green + "Portas seriais disponíveis:" + reset)
	for _, port := range services.SerialPorts {
		fmt.Println(cyan + port + reset)
	}
	fmt.Println(yellow + "-------------------------------" + reset)

	// Solicitando entrada do usuário
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow + "Escreva a porta serial em caixa alta (Ex: COM4): " + reset)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(red+"Erro ao ler a entrada:", err.Error()+reset)
		return
	}

	// Removendo espaços e nova linha do input
	input = input[:len(input)-2]
	fmt.Println(green + "Você selecionou a porta: " + cyan + input + reset)

	// Criando uma nova conexão
	connection := services.NewConnection(115200, input)

	// Inicializando a conexão
	err = connection.InitConnection()
	if err != nil {
		fmt.Println(red+"Erro ao iniciar a conexão:", err.Error()+reset)
		return
	}

	// Lendo dados da conexão
	buff := make([]byte, 5)
	connection.Read(buff)

	fmt.Println(green + "Dados lidos: " + cyan + string(buff) + reset)
}
