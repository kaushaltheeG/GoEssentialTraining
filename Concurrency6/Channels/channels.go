package main

import (
	"fmt"
	"time"
)

/*What are channels?
one direction pipe that a specfifc type
2. operations are done:  send and receive
	 receive <- [channel] <- send
If either one operation is missing, the channel will be blocked

*/

func main() {
	ch := make(chan int)

	/* gorountine w/ channel 
	//ch <- 33 //send; will be blocked since there is no one read from the channel; known as deadlock 

	//must use goroutine to overcome deadlock 
	go func() {
		ch <- 33
	}()//invokation of func
	val := <-ch //recevie 
	fmt.Printf("got %d\n", val)  
	*/

	/* Looping 
		const count = 3 
		go func() {
			for i:=0; i < count; i++ {
				fmt.Printf("sending %d\n", i)
				ch <- i //sending i to channel 
				time.Sleep(time.Second) 

			}
		}()

		//if loop is not received in a loop, it will one receive the first one 
		for i:=0; i < count; i++ {
			val := <-ch //recevie 
			fmt.Printf("got %d\n", val) 
		}

			//output
			sending 0
			got 0
			sending 1
			got 1
			sending 2
			got 2
	*/
	
	// close(ch)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i //sending i to channel 
			time.Sleep(time.Second) 
		}
		close(ch) //indicates that there will be no more values 
	}()

	for i := range ch {
			fmt.Printf("got %d\n", i) 
	}
	
	//buffer channel: can send one value without blocking 
	ch2 := make(chan int, 1) //buffered channel 
	ch2 <- 19 
	val2 := <- ch2 
	fmt.Println(val2)
}