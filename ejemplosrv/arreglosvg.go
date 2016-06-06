// arreglos

//declarar un array de 5 string cada uno, inicializar con 0
// declarar un segundo array de 5 string, inicializar con valores del primer array
// y desplegar resulyado

package main

import "fmt"

// main es la entrada del programa

func main() {
	//declarar arreglos de string que guarden los nombres
	var nombres[5]string
	
	//declarar un areglo pre-llenado con nombres de amigos
	amigos:=[5]string{"Raquel", "Isabel", "fernando", "Enrique", "jose"}

	//asignar el arreglo de los amigos al arreglo nombre
	nombres=amigos

	// mostrar cada uno de los valores string y la direccion de cada uno en la 2da

	for i, nombre := range nombres{
		fmt.Println(nombre, &nombres[i])
	}

}


