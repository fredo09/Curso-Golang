package main

import "fmt"

func main() {
	numero1 := 10
	numero2 := 20
	//mostrando resultado del funcion de la suma
	fmt.Println(sumar(numero1, numero2))

	fmt.Println("Devolviendo mas de dos valores")
	a, b := devMasValores()

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Seleccionadon un valor de los dos que se devuleven")
	_, c := devMasValores()

	fmt.Println("Imprimiendo valor de C: ", c)

	fmt.Println("Variante de una funcion")
	fmt.Println(variante1(2))

	fmt.Println("Ejecutando funcion con multiples parametros")

	fmt.Println(calculo(1, 2, 3))
	fmt.Println(calculo(1, 2, 3, 4))
	fmt.Println(calculo(1, 2, 3, 4, 5))
	fmt.Println(calculo(1, 2, 3, 4, 5, 6))
}

// funcion que suma 2 numeros
func sumar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

//funcion que devuelve mas de un valor
func vals() (int, int) {
	// devolviendo mas de un valor
	return 3, 7
}

func devMasValores() (int, bool) {
	dato := 2
	if dato < 2 {
		return dato, false
	}
	return 0, true
}

//variante de una funcion
func variante1(numero int) (z int) {
	z = numero * 2
	return
}

//funcion con parametros variables se usa "..." para indicar n numero de parametros que recibira la funcion
func calculo(numero ...int) int {
	total := 0
	fmt.Println(numero)
	// range se usa para ver el rango de una lista devuelve dos valores
	for _, num := range numero {
		total = total + num
	}

	return total
}
