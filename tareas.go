package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) marcarCompleta() {
	t.completado = true
}

func (t task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := task{
		nombre:      "Tomar desayuno",
		descripcion: "Tomar leche y comer pan",
		completado:  false,
	}
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("Almorzar")
	t.actualizarDescripcion("Plato de almuerzo con ensalada")
	fmt.Printf("%+v\n", t)
}
