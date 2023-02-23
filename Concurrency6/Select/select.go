package main

import (
	"fmt"
	"time"
)

/*
	lets us work with several channels at the same time
*/

func main()  {
	ch1, ch2 := make(chan int), make(chan int) //made two channels 

	go func() {
		ch1 <- 42 
	}()

	select { //using select statement to conditional output receive val depending on channel 
	case val := <- ch1:
		fmt.Printf("Got %d from channel 1 \n", val)
	case val := <- ch2:
		fmt.Printf("Got %d from channel 2 \n", val)
	}

	fmt.Println("----Timeout----")
	//main use case if for timeouts 
	out := make(chan float64)
	go func ()  {
		time.Sleep(10 * time.Microsecond) //adjust time to see which switch statement hits 
		out <- 3.14
	}()

	select { 
	case val := <-out:
		fmt.Printf("Got %f \n", val)
	case <- time.After(20 * time.Microsecond):
		fmt.Println("Timeout")
	}
	
}
