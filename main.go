package golang2

import "fmt"

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
}
