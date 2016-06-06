// funciones

//declaramos un struct para guardar informacion de un usuario.
// declarar la funcion que crea el valor y regresa apuntadores de este tipo y un error como valor
// hacer una segunda llamda a la funcion, pero ignorando el valor y solo probando el error como valor

package main

import "fmt"

// el usuario representa un usuario el sistemas

type usuario struct {
 nombre string
 email string
}

// nueoUsuario crea y regresa apuntadores de valore de tipo de usuario

func nuevoUsuario() (*usuario, error) {
	return &usuario{"merulo", "merulo@merulo.com"}, nil
}

// main es la ntrada de todos nuestros programas

func main(){
	// crear valores del tipo usuario

	u, err := nuevoUsuario()
	if err!= nil {
		fmt.Println(err)
		return
	}
	
	// imprimismo el valor
	fmt.Println(*u)

	// 2do llamdo a la funcion y que solo cheque el error en el regreso{return}
	_, err = nuevoUsuario()
	if err != nil {
		fmt.Println(err)
		return}
}