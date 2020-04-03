package main

import "fmt"

type task struct {
	nombre, descripcion string
	completado          bool
	prioridad           int
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
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
	var tarea2 task
	tarea2 = task{
		descripcion: "desc",
		nombre:      "nomb",
		prioridad:   1,
		completado:  true,
	}
	fmt.Println(tarea2)
}
