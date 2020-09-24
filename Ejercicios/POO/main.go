package main

import (
	"fmt"
	"time"
)

/* Herencia */
type pepe struct {
	us.usuario
}

func main() {
	u := new(pepe)
	u.altaUsuario(1, "Pepe", 12, time.Now(), true)
	fmt.Println(u.usuario)
}
