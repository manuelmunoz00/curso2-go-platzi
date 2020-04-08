package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidorC(servidor string, cnl chan string) {
	//al declarar guion bajo ignora el utilizar esta variable(contenido) ya que solo nos interesa manejar el error.
	_, err := http.Get(servidor)
	if err != nil {
		// fmt.Println(servidor, "se encuentra caído")
		cnl <- servidor + " no se encuentra disponible"
	} else {
		// fmt.Println(servidor, "OK")
		cnl <- servidor + " se encuentra OK"
	}
}

func main() {
	inicio := time.Now()
	cnl := make(chan string)
	servidores := []string{
		"https://comisariavirtual.cl",
		"https://dgd.cerofilas.gob.cl",
		"https://google.cl",
		"https://instagram.com",
	}
	//se ignora el indice utilizando guion bajo ya que solo llamaremos la funcion con el dato del slice servidor
	for _, servidor := range servidores {
		go revisarServidorC(servidor, cnl)
	}

	for index := range servidores {
		fmt.Println(<-cnl, index)
	}

	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoTranscurrido)
}
