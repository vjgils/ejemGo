// Cliente (consumidor) del chat server que consuma los datos de las goroutines y channels
// Channels y Goroutines

package main

import ( 
"io"
"log"
"net"
"os"
)

// Manejo de errores
func main() {
	conn, err := net.Dial ("tcp", "localhost:8005")
	if err != nil {
		log.Fatal(err)
	}
	// Definir variable done que ignore errores
	done := make(chan struct{})
	// Definir una función que utiliza el paquete io
	go func() {
		io.Copy(os.Stdout, conn) // ignorando errores 
		log.Println("terminamos")
		// Avisarle al goroutine principal
		done <- struct{}{}
   }()

   mustCopy(conn, os.Stdin)
   conn.Close()
   // Estamos esperando que la goroutine del background termine
   <-done
}

// Definir nueva función mustCopy
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err!= nil {
		log.Fatal(err)
	}
}