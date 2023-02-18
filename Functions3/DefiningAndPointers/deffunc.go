package main

import (
	"fmt"
)
// MUST DEFINE PARAMETER AND RETURN TYPE 
func main() {
	// val := add(1,2)
	// div, mod := divmode(1,2);
	// fmt.Println(div,mod)
	// fmt.Println(val)


	// values := []int{1,2,3,4};
	// doubleAt(values, 2)
	// fmt.Println(values);

	n := 8
	// double(n)
	// fmt.Println(n);

	doublePtr(&n) //use ambersand which passes in the variable pointer 
	fmt.Println(n)


}

//add adds a to b 
func add(a int, b int) int {
	return a +b 
}

//function can return more than one value 
func divmode(a int, b int) (int, int) {
	return a / b, a % b 
}

//get slice and index and doubles the element within slice using the given idx
func doubleAt(values []int, i int) {
	values[i] *= 2
}

//creates a new var so needs a return type 
func double(n int) {
	n *= 2
}

//pointers so passed var is being updated instead of creating a new variable in 
//pointer: *
func doublePtr(n *int) {
	*n *= 2
}