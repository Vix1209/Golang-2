package package2

import "fmt"

// Block level constants
const (
	Package2Name string = "Golang-2"
	title        string = "Package2"
	path         string = "github.com/Vix1209/Golang-2"
)

func Package2() {
	// Package level constants
	const Package2Version string = "1.0.0"
	fmt.Printf("This is the Block level constants: %s, %s, %s \n", Package2Name, title, path)
	fmt.Printf("This is the Package level constants: %s \n", Package2Version)

}
