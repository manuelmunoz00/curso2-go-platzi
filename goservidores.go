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
		//escritura en canal para retornar mensaje
		cnl <- servidor + " no se encuentra disponible"
	} else {
		// fmt.Println(servidor, "OK")
		//escritura en canal para retornar mensaje
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

	//ciclo de ejeución infinita
	i := 0
	for {
		if i = 10{
			break
		}
		for _, servidor := range servidores {
			go revisarServidorC(servidor, cnl)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-cnl)
	}

	//se ignora el indice utilizando guion bajo ya que solo llamaremos la funcion con el dato del slice servidor
	// for _, servidor := range servidores {
	// 	go revisarServidorC(servidor, cnl)
	// }

	// for index := range servidores {
	// 	//lectura del canal
	// 	fmt.Println(<-cnl, index)
	// }

	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoTranscurrido)
}
