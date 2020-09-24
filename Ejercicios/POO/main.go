package main

import (
	us "./user"
	"fmt"
)

/* Herencia */
type Pepe struct {
	us.Usuario
}

func main() {
	u := new(Pepe)
	fmt.Println(u)
}
