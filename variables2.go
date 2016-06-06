package main

import "fmt"

// variables publica
var _nuevo2 = "fsdsdfsd"
var hgdf = "fdgdf"

//camel case
var nombreUsuario = "Admin"

// fin variable publica

func main() {
	numero := 44
	_nombre := "Victor"
	numero2 := 45
	ñumero := 22

	// diferencia de mayuscula y minuscula con las variables
	Numero := 90

	fmt.Println(numero, _nombre, numero2, ñumero,
		hgdf, nombreUsuario, Numero)

	//scope
	var _nuevo = "poiupoipoi"

	fmt.Println("Prueba de varios: " + _nuevo)
	imprimir()
	imnprimir2()

}

func imprimir() {
	fmt.Println("Prueba de varios222: " + _nuevo2)
}

func imnprimir2() {
	fmt.Println(hgdf, nombreUsuario)
	fmt.Println("************************")
	hgdf, nombreUsuario = nombreUsuario, hgdf
	fmt.Println(hgdf, nombreUsuario)
}
