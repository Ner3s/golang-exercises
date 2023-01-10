package main

import (
	"fmt"
	"os"
	"os/exec"
)

func sum(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return x - y
}

func div(x, y float64) float64 {
	return x / y
}

func mul(x, y float64) float64 {
	return x * y
}

func main() {
	var running, operator int = 1, 1
	var num1, num2 float64

	for running > 0 {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("###########################")
		fmt.Println("#       CALCULADORA       #")
		fmt.Println("###########################")

		fmt.Println("Operadores: ")
		fmt.Println("1 - Somar ")
		fmt.Println("2 - Subtrair ")
		fmt.Println("3 - Multiplicar ")
		fmt.Println("4 - Dividir ")
		fmt.Print("Opção: ")
		fmt.Scan(&operator)

		fmt.Print("\nDigite primeiro número: ")
		fmt.Scan(&num1)

		fmt.Print("Digite segundo número: ")
		fmt.Scan(&num2)

		fmt.Println("\n###########################")
		fmt.Print("Resultado: ")
		switch operator {
		case 1:
			fmt.Println(sum(num1, num2))
		case 2:
			fmt.Println(sub(num1, num2))
		case 3:
			fmt.Println(div(num1, num2))
		case 4:
			fmt.Println(mul(num1, num2))
		default:
			fmt.Println("Nenhum operador selecionado!!!")
		}
		fmt.Println("###########################")

		fmt.Print("\nDeseja calcular outro valor? (1 - Sim / 0 - Não): ")
		fmt.Scan(&running)
	}

}
