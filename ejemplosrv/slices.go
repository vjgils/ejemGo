// Slices

// Declarar un slice de enteros vacío. 
// Crear un loop que meta 10 valores al slice.
// Iterar sobre el slice y mostrar cada valor.
//Declarar un slice de 5 strings e inicializar dicho slice con valores string.
//Imprimir todos los elementos.
// Tomar un slice del primer y segundo índice
// Desplegar la posición y el valor de cada uno de estos elementos en el nuevo slice.

package main

import "fmt"

// main es la entrada a cualquier programa

func main() {
	//Declarar un slice de enteros vacío
    var numeros []int

    // Meter 10 números al slice
    for i:= 0; i < 10; i++ {
     numeros = append(numeros, i*10)
    }

    // Mostrar cada valor
    for _, numero := range numeros {
    fmt.Println(numero)
    } 

    //Declarar un slice de 5 strings
    frutas := []string{"manzana", "naranja", "pera", "sandía", "aguacate"}

    // Mostrar cada índice (posición) y cada nombre de las frutas
    for i, fruta := range frutas {
    fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
    }

    // Tomar un slice de índice 1 y 2

    slice := frutas[1:3]

    //Mostrar el valor del nuevo slice
    for i, fruta := range slice {
    fmt.Printf("Index: %d Fruta: %s\n", i, fruta)
    } 

}