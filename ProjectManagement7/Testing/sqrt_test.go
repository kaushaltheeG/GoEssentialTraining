package sqrt

import (
  "testing"
  "github.com/stretchr/testify/require"
)

// func almostEqual(v1, v2 float64) bool {
// 	return Abs(v1-v2) <= 0.001
// } // //

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	
	require.NoError(t,err)
	// if err != nil {
		// 	t.Fatalf("error in calulation - %s", err)
	// }

	require.InDelta(t, 1.414214, val, 0.001) //like the almost equal func
	// if !almostEqual(val, 1.414214) {
	// 	t.Fatalf("bad value -%f", val)
	// }
}