package main

import "fmt"

func main() {
	vals := []int{1, 2, 3}

	/*
	v := vals[10] //this will cause a panic
	fmt.Println(v) //runtime error: index out of range [10] with length 3
	*/

	/* Can create custom panic but highly not recommened/ bad practice 
		panic("oops")
	*/

	output, err := safeValues(vals, 10 )
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(output)

}

//The proper way to go against panic thru the use of the defer fumc
		//will run the code cleanly 
func safeValues(vals []int, idx int) (n int, err error) {

	defer func() {
		//recover() will define e as not nil if a panic is there 
		if e := recover(); e != nil {
			//if the panic exists, the err gets formated to be the e value 
					//e get converted into an error 
			err = fmt.Errorf("%v", e)
		}
	}()

	return vals[idx], nil 
}