package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 
	} else {
		fmt.Println(ctype)
	}
}

/* 
	contentType will return the values of the Content_type header 
	returned by making an HTTP request using the url 
*/

func contentType(url string) (string, error)  {
	res, err := http.Get(url)

	if err != nil {
		return "", err
	} 

	//must closed after reading body to prevent memory leaks 
	defer res.Body.Close()

	ctype := res.Header.Get("Content-Type")

	if ctype == "" {
		//error if content type is not found 
		return "", fmt.Errorf("can't find content type")
	}

 	return ctype, nil 
}