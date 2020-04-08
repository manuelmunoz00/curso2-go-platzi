package main

import (
	"fmt"
	"net/http"
)

func revisarServidor(servidor string) {
	//al declarar guion bajo ignora el utilizar esta variable(contenido) ya que solo nos interesa manejar el error.
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, " se encuentra caído")
	} else {
		fmt.Println(servidor, " OK")
	}
}

// func main() {
// 	inicio := time.Now()
// 	servidores := []string{
// 		"https://comisariavirtual.cl",
// 		"https://dgd.cerofilas.gob.cl",
// 		"https://google.cl",
// 		"https://instagram.com",
// 	}
// 	//se ignora el indice utilizando guion bajo ya que solo llamaremos la funcion con el dato del slice servidor
// 	for _, servidor := range servidores {
// 		revisarServidor(servidor)
// 	}
// 	tiempoTranscurrido := time.Since(inicio)
// 	fmt.Printf("Tiempo de ejecucion %s\n", tiempoTranscurrido)
// }
