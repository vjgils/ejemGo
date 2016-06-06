package main

import (
"fmt"
"sync"
)

const (
goroutines = 20
capacidad = 4)

// wg se usa para esperar que el programa termine


var wg sync.WaitGroup

// recursos es un buffered channel para manejar los strings

var recursos = make(chan string, capacidad)

// main es la entrada a nuestro programa

func main() {
    // agregar el número de goroutines al waitgroup
    wg.Add(goroutines)

    // lanzar las goroutines necesarias para manejar el trabajo.
    // Asegurar de poner la llamada para saber que las gorutines terminaron
    for gr := 1; gr <= goroutines; gr++ {
        go func(gr int) {
            worker(gr)
            wg.Done()
        }(gr)
    }
    // Añadimos los strings.
    for s := 'A'; s < 'A' + capacidad; s++ {
        recursos <- string(s)
    }
    wg.Wait()
}

// lanzamos worker como una goroutine que procesa el trabajo del buffer channel

func worker(worker int){
    // Recibir un string del channel
    valor := <- recursos
    //  Imprimir el valor
    fmt.Printf("worker: %d : %s\n, ", worker, valor)
    // Poner el string de regreso
    recursos <- valor
}