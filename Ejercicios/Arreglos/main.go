package main

/*
	Arreglos estaticos y slices
*/

import "fmt"

func main() {
	fmt.Println("Arreglos estaticos y slices")

	var a [5]int

	fmt.Println("Emp:", a)

	// asiganando valores en el arreglo
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//Funcion len() se usa para ver el tamaño de un arreglo
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//Arreglos bidimensionales
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
