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

type taskList struct {
	tasks []*task
}

func (tl *taskList) agregarALista(tarea *task) {
	tl.tasks = append(tl.tasks, tarea)
}

func (tl *taskList) imprimirListaCompletados() {
	for _, tarea := range tl.tasks {
		if tarea.completado {
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripción:", tarea.descripcion)
		}
	}
}

// func main() {
// 	tarea1 := &task{
// 		nombre:      "Tomar desayuno",
// 		descripcion: "Tomar leche y comer pan",
// 		completado:  false,
// 	}
// 	// fmt.Printf("%+v\n", t)
// 	// t.marcarCompleta()
// 	// t.actualizarNombre("Almorzar")
// 	// t.actualizarDescripcion("Plato de almuerzo con ensalada")
// 	// fmt.Printf("%+v\n", t)
// 	tarea2 := &task{
// 		descripcion: "Cazuela de pollo y ensalada",
// 		nombre:      "Almorzar",
// 		prioridad:   1,
// 		completado:  false,
// 	}
// 	// fmt.Println(tarea2)

// 	tarea3 := &task{
// 		descripcion: "Fideos con salsa de tomate",
// 		nombre:      "Cenar",
// 		prioridad:   1,
// 		completado:  true,
// 	}

// 	lista := &taskList{
// 		tasks: []*task{
// 			tarea1, tarea2,
// 		},
// 	}
// 	fmt.Println(lista.tasks[0])
// 	lista.agregarALista(tarea3)
// 	fmt.Println(len(lista.tasks))

// 	for i := 0; i < len(lista.tasks); i++ {
// 		fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
// 	}

// 	for index, tarea := range lista.tasks {
// 		fmt.Println("Index", index, "nombre", tarea.nombre)
// 	}
// 	lista.tasks[0].marcarCompleta()
// 	fmt.Println("Listado de tareas completadas")
// 	lista.imprimirListaCompletados()

// 	//Se agrega mapa
// 	mapa := make(map[string]*taskList)
// 	mapa["manuel"] = lista
// 	mapa["juan"] = lista
// 	fmt.Println("Lista de tareas de Manuel")
// 	mapa["manuel"].imprimirListaCompletados()

// }
