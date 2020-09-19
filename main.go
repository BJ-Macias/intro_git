package main

import "fmt"

func main() {
	var nombre string
	fmt.Println("Hola mundo")
	fmt.Println("Nombre: ")
	fmt.Scanf("%s", &nombre)

	fmt.Println("Hola", nombre)
}