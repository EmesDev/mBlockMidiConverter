package main

import (
	"bufio"
	"fmt"
	"mBlockConverter/services"
	"os"
)

func main() {

	services.GetPorts()

	for _, port := range services.SerialPorts {
		fmt.Println(port)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Software feito por:  https://github.com/EmesDev")
	fmt.Print("Escreva a porta serial em caixa alta (Ex: COM4): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = input[:len(input)-2]
	fmt.Println("You entered:", input)

	connection := services.NewConnection(115200, input)

	err = connection.InitConnection()
	if err != nil {
		panic(err)
	}

	buff := make([]byte, 5)

	connection.Read(buff)

}
