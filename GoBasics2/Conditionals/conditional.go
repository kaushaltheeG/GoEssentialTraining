package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is big")
	} else {
		fmt.Println(("x is not that big"))
	}

	if x > 5 && x < 15 {
		fmt.Println("just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a, b := 11.0, 12.0 

	//initialize frac and used in codition all in one line 
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	//switch statement 
	n := 2
	switch n {
	case 1: 
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: 
		fmt.Println("many")
		
	}

	//naked switch statement: condition is within the case 
	switch {
	case n > 100: 
		fmt.Println("very big")
	case n > 10: 
		fmt.Println("bigger than 10")
	default: 
		fmt.Printf("%v small", n)
		
	}
}