package main

import "fmt"

func main() {
	nome := "Ivo"
	idade := 19
	versao := 1.2

	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- digite um numero:")
	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("2- digite outro numero:")
	var comando2 int
	fmt.Scan(&comando2)

	fmt.Println("O primeiro numero escolhido foi", comando)
	fmt.Println("O segundo numero escolhido foi", comando)
}
