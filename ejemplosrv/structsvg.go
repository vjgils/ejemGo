// structs

// los ejercicios:
// declara un struct para mantener info usuario (nombre, direccion,edad)

// crear un valor y lo vamos a inicializar con valores
// mostrar cada campo
// declarar e inicializar un struct anonimo con los mimso 3 campos
//mostrar el valor

package main

import "fmt"

// usuario representaa un usario en el sistemas

type usuario struct {
	nombre string
	direccion string
	edad int
}

// main es el punto de entrada de nuestra aplicaciones/sistemas

func main() {
	//depclara la variable usuario y la iniciaos usando struct
	meru := usuario{
		nombre : "merulo",
		direccion :" calle 12",
		edad : 38,
	}

	// mostrar los valores

	fmt.Println("nombre", meru.nombre)
	fmt.Println("direccion", meru.direccion)
	fmt.Println("edad", meru.edad)

	//declarar otro struct anon con los mismo 3 campos

	nicole := struct { 
		nombre string
		direccion string 
		edad int
	}{nombre: "nicole", direccion: "calle 13", edad: 22,}

	// imprimer datos
	fmt.Println("nombre", nicole.nombre)
	fmt.Println("direccion", nicole.direccion)
	fmt.Println("edad", nicole.edad)
}
