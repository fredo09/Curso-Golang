package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main(){
	leoArchivo()
	leoArchivo2()
	agregarInfo()
}

// Leer archivo en ioutil
func leoArchivo(){
	archivo, err := ioutil.ReadFile("./archivo.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

// leer archivo con OS
func leoArchivo2(){
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan(){
			registro := scanner.Text() // convertir esa linea me convierte en una cadena de texto
			fmt.Println("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func agregarInfo(){

}
