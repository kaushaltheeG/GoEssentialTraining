//unlike other processes or threads, goroutines can spin serveral of them on a single machine

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	//making a get req with given url 
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		return 
	}


	defer resp.Body.Close()

	//getting the Content-Type from response header 
	ctype := resp.Header.Get("Content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(urls []string)  {
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup //like the queue

	for _, url :=  range urls {
		wg.Add(1) //will add one to the sync

		//go key word to do a goroutine call  
		go func(url string) {
			returnType(url)
			wg.Done() //signals it is done 
		}(url)
	}
	wg.Wait() //wait for all the goroutines 
}

func main() {

	urls := []string{
		"https://golang.org",
		"https://api.github.com", 
		"https://httpbin.org/ip",
	}

	//non-concurrent way 
	start := time.Now()
	siteSerial(urls) //took 1.76 second
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrent(urls) //took 181 MS; so much FASTER 	
	fmt.Println(time.Since(start))
}

