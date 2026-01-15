package main

import (
	"fmt"
	"math/rand"
	// "github.com/Vix1209/Golang-2/golang2"
)

func Package1() {
	// Print a message to indicate that this function is called
	fmt.Println("Package1 function called")

	x := rand.Intn(300)

	fmt.Printf("Random number from package1: %d\t \n", x)

	if x <= 100 {
		fmt.Printf("The number %d is less than or equal to 100 \n", x)
	} else if x >= 101 && x <= 200 {
		fmt.Printf("The number %d is between 101 and 200 \n", x)
	} else if x >= 201 && x <= 300 {
		fmt.Printf("The number %d is between 201 and 300 \n", x)
	}

	// Call the Fruits function from the golang2 package
	// golang2.Fruits()
}

func main() {
	Package1()
}
