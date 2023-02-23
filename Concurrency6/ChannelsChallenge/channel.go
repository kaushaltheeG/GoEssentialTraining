package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string)  {
	rsp, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s -> error: %s\n", url, err)
		return 
	}

	defer rsp.Body.Close()

	ctype := rsp.Header.Get("Content-Type")
	// fmt.Printf("%s ctype is %s\n", url, ctype)
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main()  {
	urls := []string{
		"https://golang.org",
		"https://api.github.com", 
		"https://httpbin.org/ip",
	}

	//create response channel 
	ch := make(chan string) //made channel with str type 
	for _,url := range urls {
		//invoke goroutine 
		go returnType(url, ch) 
	
	}

	//blocks channel unless goroutine has sent a value via channel 
	for range urls { //Run number of URLs time 
		out := <- ch //will block until value in the channel (or go routine is done)
		fmt.Println(out) 
	}
	

}