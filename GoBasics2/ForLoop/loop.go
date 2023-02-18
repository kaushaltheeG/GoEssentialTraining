package main

import (
	"fmt"
)

func main()  {
	for i := 0; i < 3 ; i++ {
		fmt.Println(i)
	}

	fmt.Println("---Use of break---")
	//
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("---Use of Continue---")
	//go to next iteration with out executing that code
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}

		fmt.Println(i)
	}

	fmt.Println("---Go's While condition is true Loop---")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("---Go's While true Loop---")
	b := 0
	for {
		if b >2 {
			break
		}
		fmt.Println(b)
		b++
	}
	
}