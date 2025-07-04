package package2

import "fmt"

const (
	Package2Name string = "Golang-2"
	title        string = "Package2"
	path         string = "github.com/Vix1209/Golang-2"
)

func Package2() {
	fmt.Printf("This is the Package2 function constants: %s, %s, %s \n", Package2Name, title, path)
}
