package main

import "fmt"

func main() {
	var nombre string
	var edad int

	fmt.Println("Hola mundo")
	fmt.Println("Nombre: ")
	fmt.Scan(&nombre)
	fmt.Println("Hola", nombre)
	fmt.Println("Edad:")
	fmt.Scan(&edad)
	println("Vas a cumplir: ", edad + 1)


}