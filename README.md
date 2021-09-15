# Any.Go

A library to generate random values for variables and types members.

Usage:

package main

import (
	"fmt"

	"./Any"
)

func main() {
	fmt.Println("Random Country: ", Any.Country())
	fmt.Println("Random Month: ", Any.Month())
	fmt.Println("Random Month: ", Any.UsState())
	fmt.Println("Random Month: ", Any.SolarSystemPlanet())
	fmt.Println("Random Month: ", Any.DayOftheWeek())
}

