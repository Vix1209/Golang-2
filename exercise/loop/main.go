package main

import "fmt"

func testRangeLoop() {
	// range loop
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	// fmt.Printf("Map contents: %v\n", m)
	for k, v := range m {
		fmt.Printf("Key: %s \t Value: %d\n", k, v)
	}
}

func main() {
	// for i := range 100 {
	// 	fmt.Println(i)
	// }

	testRangeLoop()
}
