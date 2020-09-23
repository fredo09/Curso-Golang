/*
* Formas de crear un mapa en GOLANG
 */

package main

import "fmt"

func main() {
	/*
	* Formas de crear un mapa en GOLANG
	 */

	paises := make(map[string]string, 5)
	fmt.Println(paises)

	//primer ejemplo de mapas
	paises["Mexico"] = "Cdmx"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	// segundo ejemplo de mapas
	campeonatoFutbol := map[string]int{
		"Barcelona":   50,
		"Real Madril": 50,
		"Chivas":      100}

	fmt.Println(campeonatoFutbol)

	//agregando o modificando  y borrar valores a un mapa
	campeonatoFutbol["Chelese"] = 100

	delete(campeonatoFutbol, "Real Madril")

	fmt.Println(campeonatoFutbol)

	for equipo, puntaje := range campeonatoFutbol {
		fmt.Println("El equipo , tiene un puntaje de: \n ", equipo, puntaje)
	}

	//ver si un elemento existe en un mapa
	puntaje, existe := campeonatoFutbol["Tigres"]
	fmt.Printf("El puntaje capturado es %d, y el quipo exite %t \n", puntaje, existe)
}
