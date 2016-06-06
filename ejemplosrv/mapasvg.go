// mapas

// declarar y hacer un mapa de valores enteros con un string como llave.
// llenar el mapa con 5 valores e iterar sobre el mapa para mostrar los pares llave/valor

package main 

import "fmt"

// main entrada del programa fijo

func main() {
	// declarar y hacer mapas
	departamentos := make(map[string]int)

	// inicializar los datos en el mapa
	departamentos["devs"]=25
	departamentos["marketing"]=50
	departamentos["ejecutivos"]=4
	departamentos["ventas"]=60
	departamentos["mantenimiento"]=8

	// desplegar por medio de iteracion el valor de cada par llave/valor
	for key, value := range departamentos {
		fmt.Printf("Dept: %s Personas: %d\n", key, value)
	}
}