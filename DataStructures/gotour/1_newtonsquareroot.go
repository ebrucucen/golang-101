package newtonsquareroot

import (
	"fmt"
	"math"
)

func newtonsquareroot(){
	var x,z float64=64,64
	libsquare(x)
	newtonsquare(x, z)
}
func libsquare(x float64) float64 {
	r := math.Sqrt(x)
	fmt.Printf("result from lib %f\n", r)
	return r
}
func newtonsquare(x, z float64) float64 {

	fmt.Printf("Value: %f\n", z)
	for count := 1; count < 10; count++ {
		z = z - (((z * z) - x) / 2 / z)
		fmt.Printf("Value: %f\n", z)
	}
	return z
}
