/* 
	You need to print the numbers between 1 and 20, one per line. 
		- if the number is divisible by 3, let's say 9, you need to print fizz instead of the number. 
		- if the number is visible by 5, let's say 10, you need to print buzz instead of the number. 
		- if the number is visible by both three and five, let's say 15, you need to print fizz buzz. 
		- And otherwise you just print the number. 
*/


package main

import (
	"fmt"
)


func main()  {
	
	for i := 1; i < 21; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0: 
			fmt.Println("fizz buzz")
		case i % 3 == 0:
			fmt.Println("fizz")
		case i % 5 == 0: 
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	} 
}
	

