package main

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	return 1
}

/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	operador := "-"
	valores := strings.Split(operacion, operador)
	valor1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println("Hubo un error en el primer argumento")
	}
	valor2, err2 := strconv.Atoi(valores[1])
	if err2 != nil {
		fmt.Println("Hubo un error en el segundo argumento")
	}

	switch operador {
	case "+":
		fmt.Println(valor1 + valor2)
	case "-":
		fmt.Println(valor1 - valor2)
	case "*":
		fmt.Println(valor1 * valor2)
	case "/":
		fmt.Println(valor1 / valor2)
	default:
		fmt.Println(operador, "no esta soportado")
	}

}
*/
