// init package
package main

// immport modulos / libs
import "fmt"

// Atribuição de variaveis

var num1 int
var num2 int
var operator string

func main() {

	fmt.Println("digite o primeiro número:")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&num2)
	fmt.Println("Digite a operação:")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Println("Resultado:", num1+num2)
	case "-":
		fmt.Println("Resultado:", num1-num2)
	case "*":
		fmt.Println("Resultado:", num1*num2)
	case "/":
		fmt.Println("Resultado:", num1/num2)

	}
}
