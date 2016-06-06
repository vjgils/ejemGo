// slices

// declarar un slices de enteros vacio
// crear un loop que meta 10 valores al slices

// iterar sobre el slice y mostrar cada valor

// decalrar un slices de 5 string e inicializar con valores string
// imprimir todos los elementos
// tomar un slices del primer y segundo indice
// desplegar la posicion y el valor de cad uno de estos elementos en el nuevo slice.

package main

import "fmt"

// main es la entrada de cualquier programa

func main() {
	// declarar un slice de entero vacios
	varf numeros []int

	// meter numeros al slices
	for i:= 0; i<=10 i++ {
		numeros = append(numeros, 1*10)
	}

	//mostrar cada valor
	for _, numero := range numeros{
		fmt.Println(numero)
	}

	// declarar un slice de cadena

	frutas:= []string{"manzanas","naranjas","pera", "sandias", "cambur"}

	// mostrar cada indice (posicion) y cada nombre de las frutas
	for i, fruta:= range frutas{
		fmt.Printf("Index: %d fruta: %s\n", i, fruta)

	}
	 // tomar un slices de indice 1 y 2
	slice := frutas[0:3]

	// mostrar el valor del nuevo slice
	for i, fruta := range slice{
		fmt.Printf("Index: %d fruta: %s\n", i, fruta)
	}
}