// Mapas

// Declarar y hacer un mapa de valores enteros con un string como llave.
// Llenar el mapa con 5 valores 
// Iterar sobre el mapa para mostrar los pares llave/valor

package main

import "fmt"

// main es la entrada de nuestro programa

func main() {
	//Declarar y hacer el mapa

	departamentos := make(map[string]int)

	// Inicializar los datos en el mapa

	departamentos["Devs"] = 25
	departamentos["Marketing"] = 50
	departamentos["Ejecutivos"]= 4
	departamentos["Ventas"]= 60
	departamentos["Mantenimiento"] = 8

	// Desplegar por medio de iteración el valor de cada par llave/valor
for key, value := range departamentos {
	fmt.Printf("Dept: %s Personas: %d\n", key, value)
}
}