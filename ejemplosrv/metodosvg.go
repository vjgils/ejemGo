// metodos

// declarar un struct que representa jugador (nombre goles tarjetas)
// declarar un metodo que calcule el promedio de goles de un jugador
// partidos/goles

// declarar slices que inicialize un slice con varios jugadores
// iterar sobre el slice mostrando los jugadores y sus promedios de goles y chutes

package main 

import "fmt"

// jugador represneta una persona

type jugador struct{nombre string, goles int, partidos int}

// promedio de goles

func (j *jugador) average()float64{
	return float64(j.partidos)/ float64(j.goles)
}

// main es la entradad

func main() {
	// crear algunos jugadores

	jugaores:=[]jugador {
		{"carlos", 20, 60}
		{"fernando", 17, 30}
		{"alonso",34,60}
	}

	// mostrar el promedio de goles para cada jugador
	for i := 0; i < count; i++ {
		
	}
}