package main

import (
	"fmt"
)

func main() {
	var age int
	var name string

	fmt.Print("Bitte gebe dein Alter ein: ")
	fmt.Scan(&age)
	fmt.Print("Wie heißt du? ")
	fmt.Scan(&name)
	fmt.Printf("Hallo %v, du bist %v Jahre alt!\n", name, age)

	if age < 18 {
		var missing = 18 - age
		var ages = "Jahre"
		if missing == 1 {
			ages = "Jahr"
		}
		fmt.Printf("Du bist NICHT volljährig. Du brauchst %v %v, bis du volljährig bist.", missing, ages)
	} else {
		fmt.Println("Du bist volljährig.")
	}
}
