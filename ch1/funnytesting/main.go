//to test basics
package main

import (
	"fmt"
	"math"
)

func main() {

	x := math.Mod(2, 2)
	y := math.Mod(3, 2)
	fmt.Printf("x: %d\n", int(x))
	fmt.Printf("y: %d\n", int(y))

	const (
		cycles  = 5     //number of complete x iscullator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers [-size ..+size]
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames in 10 ms units
	)
	//freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	//phase := 0.0                 //phase difference
	fmt.Printf("pi:%f", 10*math.Pi)
	/*for t := 0.0; t < 10*math.Pi; t += res {
	x := math.Sin(t)
	fmt.Printf("x:%f\n", x)
	fmt.Printf("t:%f\n", t)

	x = math.Sin(t)
	y = math.Sin(t*freq + phase)
	//		index := math.Mod(nframes, 4)
	fmt.Printf("a:%d\n", (size + int(x*size+0.5)))
	fmt.Printf("b:%d\n", (size + int(y*size+0.5)))
	}*/
}
