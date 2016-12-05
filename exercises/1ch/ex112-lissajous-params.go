// Acts like a Lissajous figure web server. The parameters controlling the form of figures
// can be passed on the URL (in the form of the query string parameters).

package main

import (
	"image/gif"
	"image"
	"math"
	"math/rand"
	"io"
	"image/color"
	"net/http"
	"log"
	"strconv"
	"fmt"
)

var palette = []color.Color{color.Black, color.RGBA{}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	}) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func lissajous(out io.Writer, r *http.Request) {
	var cycles int = 5     // number of complete x oscillator revolutions
	var res  float64   = 0.001 // angular resolution
	var size  int  = 100   // image canvas covers [-size..+size]
	var nframes int = 64    // number of animation frames
	var delay   int = 8     // delay between frames in 10ms units
	for k, v := range r.URL.Query() {
		fmt.Printf("key: %v \t val: %v\n", k, v)
		if k == "cycles" {
			x, err := strconv.Atoi(v[0])
			if err == nil {
				cycles = x
			}
		} else if k == "res" {
			x, err := strconv.ParseFloat(v[0], 64)
			if err == nil {
				res = x
			}
		} else if k == "size" {
			x, err := strconv.Atoi(v[0])
			if err == nil {
				size = x
			}
		}
	}
	fmt.Printf("cycles: %d\n", cycles)
	fmt.Printf("res: %f\n", res)
	fmt.Printf("size: %d\n", size)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
