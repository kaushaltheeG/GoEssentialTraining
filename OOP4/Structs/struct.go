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
	//If Field is Captilized, it is accessable outside of the package (public) aka exported symbols,
	//else if lower case, the field is only accessable within the package (private) aka unexported symbols

	CampaignID string 
	Balance float64 //USD
	Expires time.Time 
}

//defining a methods for budget 
// b ensures that is method belong only for an instance of Budget 
//TimeLeft is the method name & time.Duration is the return value 
func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC()) // take in expiration date and subtracts current time in UTC 
}

// * (pointer receiver which will change that instance field's balance ) 
func (b *Budget) Update(sum float64) {
	b.Balance += sum
}


func main() {
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1.TimeLeft())

	b1.Update(10.5)
	fmt.Println(b1.Balance)

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
