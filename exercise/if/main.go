package main

import (
	p "exercise/if/printer"
	"fmt"
	"math/rand"
)

func importedPrint() {

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("these are the values of the two integers, x and y: %d and %d\n", x, y)

	if x < 4 && y < 4 {
		p.Print("both x and y are less than 4", p.Line)
	}

}

func commaOkIdiom(x string) {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	// the comma ok idiom in Go. confirms the existence of something then runs conditions based on that.
	if v, ok := m[x]; !ok {
		fmt.Printf("this key does not exist in the map, %v\n", v)
	} else {
		fmt.Printf("this key exists in the map and its value is %d\n", v)
	}
}

func statementStatementIdion() {

	for key := range 100 {
		fmt.Printf("key is %d \t", key+1)
		if x := rand.Intn(5); x < 3 {
			fmt.Printf("x is less than 3: %d\n", x)
		} else {
			fmt.Printf("x is greater than or equal to 3: %d\n", x)
		}
	}

	fmt.Println("The benefit of using the statement statement idiom is that for every iteration of the loop, a new random number is generated and evaluated within the if statement.")

}

func checkDifferentPrintStatements() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func main() {

	// importedPrint()

	// commaOkIdiom("Ofor")

	// statementStatementIdion()

	checkDifferentPrintStatements()
}
