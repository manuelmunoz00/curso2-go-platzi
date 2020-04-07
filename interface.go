package main

import "fmt"

type perro struct {
}

type pescado struct {
}

type ave struct {
}

//Los structs anteriores deben implementar el metodo mover para cada uno que es requisito para la interfaz de tipo animal
func (p perro) mover() string {
	return "Soy un perro y estoy caminando"
}

func (p pescado) mover() string {
	return "Sour un pescado y estoy nadando"
}

func (p ave) mover() string {
	return "Soy un ave y estoy volando"
}

//Interfaz de tipo animal
type animal interface {
	mover() string
}

//Función que recibe un animal
func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

// func main() {
// 	pr := perro{}
// 	moverAnimal(pr)
// 	pc := pescado{}
// 	moverAnimal(pc)
// 	av := ave{}
// 	moverAnimal(av)
// }
