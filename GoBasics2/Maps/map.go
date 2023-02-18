//map data structure 

package main

import (
	"fmt"
	"strings"
)

func main()  {
	stocks := map[string]float64 { //key is str and val is float64
		"AMZN": 2087.98,
		"GOOG": 2540.87,
		"MSFT": 273.89, //must have trailing comma in multi line 
	}

	//numbers of items 
	fmt.Println(len(stocks)) //prints amt of keys 

	//keying into Map
	fmt.Println(stocks["MSFT"])

	//using a key thjat does not exists: get value of 0
	fmt.Println(stocks["TSLA"])

	//using two values to see if key exists
	value, ok := stocks["TSLA"];

	if !ok {
		fmt.Println("Stocks not found")
	} else {
		fmt.Println(value)
	}

	//Setting a new key,value pair 
	stocks["TSLA"] = 822.22

	//Deleting from map
	delete(stocks, "AMZN")

	//looping using keysthru Map
	for key := range stocks {
		fmt.Println(key)
	}

	//loop with key and value
	for key, value := range stocks {
		fmt.Printf("key = %s and value = %.2f \n", key, value)
	}

	//Challenge 
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail 
	To catch me the wind
	`

	words := strings.Fields(text) //splits string by spaces and adds to an array; similar to split(' ') in js
	fmt.Println(words)
	counts := map[string]int{}
	for _,word := range words {
		//if it does not exist, it returns 0 so it adding and incrementing the value 
		counts[strings.ToLower((word))]++
	}

	fmt.Println(counts)


}