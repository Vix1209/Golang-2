package package1

import (
	"fmt"

	golang2 "github.com/Vix1209/Golang-2"
)

func Package1() {
	// Print a message to indicate that this function is called
	fmt.Println("Package1 function called")

	// Call the Fruits function from the golang2 package
	golang2.Fruits()
}
