//Print the max value in the slice 

package main

import (
	"fmt"
)

func main()  {
	
	nums := []int {16, 8, 42, 4, 23, 15};
	fmt.Println(nums)

	max := 0;

	for _,num := range nums {
		if max < num {
			max = num 
		}
	}
	println(max)
}