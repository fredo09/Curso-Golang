package main

import "fmt"

func main() {
	var edad int
	edad = 17

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("No eres mayor de edad")
	}
}
