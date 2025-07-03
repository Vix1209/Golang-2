package golang2

import (
	"fmt"

	golang3 "github.com/Vix1209/Golang-3"
)

func Fruits() {

	// Define a map with string keys and integer values
	myMap := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	// Iterate over the map using a for range loop
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	fmt.Println("Iteration complete.")

	golang3.DirectDependency()
}
