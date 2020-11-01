package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func parsear(entrada string) (int, error) {
	valor, err := strconv.Atoi(entrada) // strcon.Atoi string a number
	if err != nil {
		return 0, err
	}

	return valor, nil
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func (calc) operate(entrada string, operador string) int {
	var result int
	entradaLimpia := strings.Split(entrada, operador)
	valor1, err := parsear(entradaLimpia[0])
	valor2, err := parsear(entradaLimpia[1])

	if err != nil {
		fmt.Println(err)
	}
	switch operador {
	case "+":
		result = valor1 + valor2
	case "-":
		result = valor1 - valor2
	case "*":
		result = valor1 * valor2
	case "/":
		result = valor1 / valor2
	}

	return result
}

func getOperator(cadena string) string {
	var operador string

	runes := []rune(cadena)

	for _, character := range runes {
		str := string(character)
		switch str {
		case "+":
			operador = "+"
		case "-":
			operador = "-"
		case "*":
			operador = "*"
		case "/":
			operador = "/"
		}
	}

	return operador
}

func main() {
	entrada := readInput()
	operador := getOperator(entrada)
	c := calc{}
	result := c.operate(entrada, operador)
	fmt.Println(result)
}
