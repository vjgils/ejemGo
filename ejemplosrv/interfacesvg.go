// interfaces

// declarar una interface laada speaker con un metodo llamado speak
//declara un struct llamdo ingles que reprente una persona que hable ingles
//declarar un struct llamado chino que represente a una persona que hable chino
// implemetar la interface speaker para cada struct usando un valor y string: "hello world y @@@@"
// declarar una variable del tipo speaker y asignarle una direccio  de tipo ingles y llamar el metodo
// hay que hacerlo para chino

package main

import "fmt"

// speaker implementa

type speaker interface {
	speak()
}

// ingles represneta una persona que habla ingles

type ingles struct{}

// speak implementa la interface speaker 
func (ingles) speak() {
	fmt.Println("hello world")
}

// chino representa a una persona que habla chino

type chino struct{}
//speak implementa la interface speaker
func(chino) speak() {
	fmt.Println("@@@@@@")
}

func main() {
	
}
