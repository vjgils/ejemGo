// Arreglos

// Declarar un array de 5 strings cada uno, los vamos a inicializar con 0.
// Declarar un segundo arreglo de 5 strings, los vamos a inicializar con los valores de strings.
// Asignar el segundo arreglo al primero y desplegar los resultados del primero.
// Mostrar el valor de la cadena y la dirección de cada elemento.

package main

import "fmt"

// main es la entrada del programa

func main() {
	//Declarar arreglos de strings que guarden los nombres
	var nombres[5]string

	// Declarar un arreglo pre-llenado con nombres de amigos.
	 amigos:= [5]string{"Raquel", "Isabel", "Fernando", "Enrique", "José"}

	 // Asignar el arreglo de los amigos al arreglo nombres

	 nombres = amigos

	 // Mostrar cada uno de los valores string y la dirección de cada uno en el 2ndo arreglo
for i, nombre := range nombres {
       fmt.Println(nombre, &nombres[i])
	 }
}