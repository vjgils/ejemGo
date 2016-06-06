// Structs

// Los ejercicios:
// Declarar un struct para mantener info. de un usuario (nombre, dirección, edad)

// Crear un valor y lo vamos a inicializar con valores
// Mostrar cada campo
// Declarar e inicializar un struct anónimo con los mismos 3 campos
// Mostrar el valor.

package main

import "fmt"

// usuario representa un usuario en el sistema

type usuario struct {
	nombre string
	direccion string
	edad int
}

// main es el punto de entrada de nuestras aplicaciones/ sistemas

func main() {
	// Declarar la variable usuario y la iniciamos usando un struct

	vero := usuario{
	nombre : "Verónica",
	direccion : "Calle 12",
	edad : 38,
	}

	//Mostrar los valores de cada campo

	fmt.Println("Nombre", vero.nombre)
	fmt.Println("Dirección", vero.direccion)
	fmt.Println("Age", vero.edad)

	// Declarar otro struct anon con los mismos 3 campos
    nicole := struct {
          nombre string
          direccion string
          edad int
    } {

    nombre: "Nicole",
    direccion : "Calle 13",
    edad: 22,
    }

    // Imprimir datos de Nicole

    fmt.Println("Nombre", nicole.nombre)
    fmt.Println("Dirección", nicole.direccion)
    fmt.Println("Edad", nicole.edad)
}