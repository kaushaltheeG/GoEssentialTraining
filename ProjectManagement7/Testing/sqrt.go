package sqrt

import (
	"errors"
	"math"
)

//common errors
var (
	ErrNegSqrt = errors.New("sqrt of neagtive number")
	ErrNoSolution = errors.New("no solution found")
)

//abs restusn the absolute value of val 
func Abs(val float64) float64  {
	if val < 0 {
		return -val
	}
	return val 
}

//sqrt function 
func Sqrt(num float64) (float64,error) {
	return math.Sqrt(Abs(num)), nil 
}
