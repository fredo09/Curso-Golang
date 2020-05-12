package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese el numero 1")
	//fmt.Scanf("%d", &numero1) // leyendo desde la consola el numero ingresado
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el numero 2")
	//fmt.Scanf("%d", &numero2)
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	//Leemos la entrada del teclado usando el paquete OS y leyendo con BUFIO
	scanner := bufio.NewScanner(os.Stdin)

	//Controlando si hay entrada de teclado
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2

	// immprimiendo el resultado
	fmt.Println(leyenda, resultado)
}
