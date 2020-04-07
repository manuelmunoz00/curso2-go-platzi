package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorweb struct{}

func (escritorweb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	respuesta, error := http.Get("https://google.cl")
	if error != nil {
		fmt.Println("error al obtener el contenido", error)
	}
	e := escritorweb{}
	io.Copy(e, respuesta.Body)
}
