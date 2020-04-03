package main

import "fmt"

// func main() {
// 	x := 25
// 	fmt.Println(x)
// 	// cambiarValor(x)
// 	fmt.Println(&x)
// 	y := &x
// 	// cambiarValor(y)
// 	fmt.Println(y)
// 	fmt.Println(*y)
// 	// cambiarValor(*y)
// }

func cambiarValor(valor int) {
	fmt.Println(&valor)
	valor = 36
}
