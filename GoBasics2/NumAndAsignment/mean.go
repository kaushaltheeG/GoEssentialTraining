package main

import (
	"fmt"
)

func main() {
	/*  //Int: 10 different types: int,uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64
	var x int 
	var y int 

	x = 1
	y = 2
	*/


	/* // Float: 2 types: float 32, and float 64
	var x float64
	var y float64

	x = 1.0
	y = 2.0 
	*/
	// := defines and assigns the variable at the same time; no need for two line var declaration   
	// x := 1.0 
	// y := 2.0 

	//one lines for two vars declaration 
	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, x) //%v is value and %T is type of x and x 
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0 
	fmt.Printf("result: %v, type of %T/n", mean, mean)
}
