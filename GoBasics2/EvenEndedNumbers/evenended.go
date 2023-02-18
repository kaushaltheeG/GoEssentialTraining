/* 
	Q: Count how many even-ended numbers are there that are a multiplication of two four-digit numbers 

	Even-ended number: a number with the same first and last digits

*/

package main

import (
	"fmt"
)
func main()  {
	
	count := 0

	for a := 1000; a < 9999; a++ {
		for b := a; b < 9999; b++ {
			n := a * b 

			//if a * b is even ended 
			s := fmt.Sprintf("%d", n) //converts product into string

			//checks if first and last num match
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}
	fmt.Println(count)
}