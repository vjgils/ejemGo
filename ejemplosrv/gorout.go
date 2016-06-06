// Goroutines

//Crear un programa que declare 2 funciones anónimas.
// La primera de estas funciones va a contar del 100 al 0 .
// Mostrar cada número con un identificador único para cada goroutine 
// Crear goruoutines a partir de estas funciones y no permitir que main() regrese hasta que todas las goroutines se completen.

// Correr el programa en parelelo

package main

import (
 "fmt"
 "runtime"
 "sync")

// init se ejecuta antes de main()
func init() {
	// Alojando/reservando (allocate)un procesador lógico para que lo use el scheduler.
	runtime.GOMAXPROCS(1)
}

// main es la entrada a nuestros programas

func main() {
	
// Declarar un wait group (grupo de espera) para iniciar el conteo en las 2 goroutines.

var wg sync.WaitGroup
wg.Add(2)

fmt.Println("Iniciar Goroutines")

// Declarar una función anónima y crear una goroutine

go func() {
	//Cuenta regresiva del 100 al 0
	for count := 100; count >= 0; count-- {
	fmt.Printf("[A:%d]\n", count)
	}

	//Avisarle a main que ya terminamos 
	wg.Done()
}()

// Declarar una función anónima y crear una goroutine

go func(){
	//Contar del 0 al 100
	for count := 0; count<100; count++{
		fmt.Printf("[B:%d]\n", count)
	}

	//Decirle a main que ya terminamos
	wg.Done()
}()

// Esperar a que terminen las goroutines
fmt.Println("Esperando a que terminen")
wg.Wait()

fmt.Println("\nFinalizar el programa")

}








