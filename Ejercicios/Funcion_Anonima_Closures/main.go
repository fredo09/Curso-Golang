package main

import "fmt"

// asignando una funcion a una variable
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Funcion Main")

	//ejemplo de funcion anonima function suma
	fmt.Println(calculo(5, 2))

	//redefinido funcion calculo
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	// funcion resta
	fmt.Println(calculo(5, 2))

	//otro ejemplo de funcion anonima
	operaciones()

	//llamanndo a closures
	tabla1 := 2
	mitabla := tabla(tabla1)

	fmt.Println("closures")

	for i := 1; i < 11; i++ {
		fmt.Println(mitabla())
	}
}

func operaciones() {
	// otra manera de crear funcion anonima
	resultado := func() int {
		var a int = 23
		var b int = 25
		return a + b
	}

	fmt.Println(resultado())
}

//closures
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
