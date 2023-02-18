package main

import (
	"fmt"
	"math"
)

func main() {
	vaild, err := sqrt(-3.0)
	// failed, err := sqrt(-1.0)
	// fmt.Println(vaild, error)

	if err != nil {
		fmt.Printf("Errorr %s\n", err)
	} else {
		fmt.Println(vaild)
	}

}

func sqrt(n float64) (float64, error)  {
	if (n < 0) {
		//customary to return 0 if there is an error 
		return 0.0, fmt.Errorf("sqrt of a negative value (%f)", n)
	}

	//return nill for error 
	return math.Sqrt(n), nil 
	
}