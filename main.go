package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// checa se o número de argumentos é o esperado
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <number1> <operator> <number2>")
		return
	}

	// converte os argumentos para os tipos corretos
	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator:= os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	// checa se houve erro na conversão
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid number")
		return
	}

	
