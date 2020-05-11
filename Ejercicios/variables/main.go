package main

import "fmt"

// tipos de vatiables, esta con Mayusculas sera exportado a otros paquetes
var numero int
var texto string
var status bool

func main() {
	numero2, numero3, numero4 := 5, 10, 67 // forma de asiganar datos a las vatiables

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)

}
