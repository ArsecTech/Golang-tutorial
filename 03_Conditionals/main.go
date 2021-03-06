package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 { // true&&true==(true)
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 { //true || false = (true)
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	// frac = 0.55
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	switchStatements()

}

func switchStatements() {

	x := 2

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

}
