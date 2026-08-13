package main

import "fmt"

func main() {

	//* Variables mutables ( var, := )

	var name string = "Arturo"
	edad := 22
	city := "Ventanilla"
	height := float64(1.71)
	learningGolang := true

	edad = 23
	city = "Puente Piedra"

	var country string

	fmt.Println(name)
	fmt.Println(edad)
	fmt.Println(city)
	fmt.Println(height)
	fmt.Println(learningGolang)
	fmt.Println(country)

	//* Variables inmutables ( const )

	const countryV2 = "Peru"
	const daysOfWeek = 7

	fmt.Println(countryV2)
	fmt.Println(daysOfWeek)

	// countryV2 = "Onichan"
}
