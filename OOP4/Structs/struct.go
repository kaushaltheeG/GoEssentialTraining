package main

import (
	"fmt"
	"time"
)

/* 
	Struct: the class that allows you to keep track of several fields
*/

type Budget struct {
	//below is the fields below within it associated types 
	//If Field is Captilized, it is accessable outside of the package (public),
	//else if lower case, the field is only accessable within the package (private)
	
	CampaignID string 
	Balance float64 //USD
	Expries time.Time 
}


func main() {
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)

	fmt.Printf("%#v\n", b1) //allows you to see more; such as field name 

	//keying into a field 
	fmt.Println(b1.CampaignID)

	//defining a struct while omiting some fields 
	//Omited fields are given there default value of 0; for time it is jan 1
	b2 := Budget{
		Balance: 19.3,
		CampaignID: "pups",
	}

	fmt.Printf("%#v\n", b2)

	var b3 Budget
	fmt.Printf("%#v\n",b3)


}
