package main

import (
	"fmt"
)

func main() {
	//STRINGS ARE IMUTABLE LIKE JAVASCRIPT 
	book := "The colour of magic"
	fmt.Println(book)

	//Length function len()
		//returns amount of bytes in the string 
	fmt.Println(len(book))

	//keying into the first byte with square brackets 
	fmt.Printf("book[0] = %v (type of %T)\n", book[0], book[0]) //prints 84 which is number for B and uint8 which in Go stands for byte 

	//will get error bc it is imutable 
	// book[0] = 116

	//Slice (start to end) a part of the string; end is exculed 
		//aka Half empty range 
	fmt.Println(book[4:11])

	//Slice (no end): from start to end 
	fmt.Println(book[4:])

	//Slice (no start) from 0 up to but not including the end index 
	fmt.Println(book[:4])

	//Concatenate strings
	fmt.Println("y" + book[1:2])

	//Unicode: can write like a half symbol within a string
	fmt.Println("It was 1/2 price")

	//Multi Line by using raw strings
		//raw strings: uses backtick, which they ignore the back slash special characters like /N and /T so you can write multi line strings
	poem := `
		The road goes ever on
		Down from the door where it began
		...
	`
	fmt.Println(poem)

}