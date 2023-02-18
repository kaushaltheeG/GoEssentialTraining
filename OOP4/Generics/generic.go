package main

import "fmt"

/*
	Generics: an interface can also be a set of types
*/

//this ordered interface can be an int, float64, or a string
type Ordered interface {
	int | float64 | string
}

//this min function can now accept either an int, float64, or string slice  
func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("min of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m  {
			m = v 
		}
	}
	return m, nil 
}

func main()  {
	fmt.Println(min([]float64{2,1,3}))
	fmt.Println(min([]string{"c", "b", "c"}))

}