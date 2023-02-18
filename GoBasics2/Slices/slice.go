/* 
	Slice: is a sequence of items equivalent to a list or array in another language 
*/

package main

import (
	"fmt"
)

func main()  {
	//Same type 
	loons := []string{"bugs", "daffy", "taaz"};
	fmt.Printf("loons = %v (type of %T)\n", loons, loons)

	//length 
	fmt.Println(len(loons))

	//key into 
	fmt.Printf("loon[0] = %v \n", loons[0])

	//slicing (no end)
	fmt.Println(loons[1:])

	//traditional loop 
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	//Range loop: from 0 to n 
	for i := range loons {
		fmt.Println(loons[i])
	}

	//Range loop with two args: i is idx and name is the element in the slice(array)
	for i,name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	//Range but only elements No idx: use the _ == undefined 
	for _,name := range loons {
		fmt.Printf("%s \n", name)
	}

	//appending to the end of the slice 
	loons = append(loons, "elmer") 
	fmt.Println(loons)
}