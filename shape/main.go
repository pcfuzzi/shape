package main

import (
	"fmt"
)

func main() {
	// Erstellen eines neuen Rechtecks
	r := Rectangle{Width: 10, Height: 5}

	// Aufrufen der Methoden des Rechtecks

	fmt.Println("Fläche:", r.Area())
	fmt.Println("Umfang:", r.Perimeter())
}
