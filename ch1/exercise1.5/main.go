//Change Lissajous' program color palette to green on black
//create web color #RRGGBB color.RGBA{0xRR,0xGG,0xBB,0xff}
//Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 1}, color.RGBA{0x00, 0xFF, 0x00, 1}}

const (
	whiteIndex = 0 //first color in palette
	blackIndex = 1 //next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //number of complete x iscullator revolutions
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers [-size ..+size]
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames in 10 ms units
	)
	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTE: ignore errors
}

/*
bash-3.2$ go build
bash-3.2$ ./main >out.gif
*/
