package main

import (
	"fmt"
)

/* 
	Even though Go has a Garbage Collector to remove unused objects, 
	there other resources that Go may use like Virtual Machine, Sockets, Etc
	that must be safely opened and closed after use

	Defer is key word used to run a resource clean up function after the desired 
	function is done running 
*/

func main()  {
	worker() // worker 
			// Clean up b 
			//clean up a 
	
}

func worker()  {
	r1, err := acquire("a")
	if err != nil {
		fmt.Println("ERROR", err)
		return 
	}

	//this function is called after the function has finsihed running; will release the resource after use
	defer release(r1)

	r2, err := acquire("b")
	if err != nil {
		fmt.Println("ERROR:", err)
		return 
	}
	defer release(r2)

	fmt.Println("worker")
}

//getting the "resource"
func acquire(name string) (string, error)  {
	return name, nil 
}

//used with defer to close the resource
func release(name string) {
	fmt.Printf("Clean up %s\n", name)
}