package main

import (
	"fmt"
	"math/rand"
)

func anotherSwitch() {
	xi := []int{1, 2, 3, 4, 5} // slice of integers, this is a type of array but more flexible. because its size can grow and shrink. how?? because its built on top of an array and it manages the size internally.
	for i, v := range xi {
		x := rand.Intn(5)
		switch {
		case x < 2:
			fmt.Printf("%d less than 2, \n", x)
		case x >= 2:
			fmt.Printf("%d greater than or equal to 2, \n", x)
		default:
			fmt.Printf("%d everywhere good, \n", x)
		}
		fmt.Printf("iteration number: %d \t value is %d \n", i, v)
		fmt.Println("")
	}
}

func anotherSwitchAgain() {
	const x int = 40

	switch {
	case x < 30:
		println("this is less than 30")
		// fallthrough

	case x == 40:
		println("this is equal to 40")
	default:
		println("everywhere good")
	}
}

func main() {
	// anotherSwitch()

	anotherSwitchAgain()
}
