package main

import (
	"bufio"
	"fmt"
	"image"
	"math/cmplx"
	"image/png"
	"os"
)

var depth uint
const MAXDEPTH = 6
var M int
const MAXM = 40000

const c_left_top = complex(-2, 2)
const c_right_bottom = complex(2, -2)
const NMAX = 100

var complete chan int
var step float64
var rlen float64
var ilen float64

func main() {
	fmt.Sscan(os.Args[1], &M)
	if (M > MAXM) {
		fmt.Printf("The picture size %dx%d is too big. Exit.\n", M, M)
		return
	}
	fmt.Sscan(os.Args[2], &depth)	
	if (depth > MAXDEPTH) {
		fmt.Printf("The number of threads(%d) is too big. Exit.\n", 1 << depth)
		return
	}

	rlen = real(c_right_bottom) - real(c_left_top)
	ilen = imag(c_right_bottom) - imag(c_left_top)
	step = rlen / (float64)(M)

	Mandelbrot()
}

func Mandelbrot() {
	r := image.Rect(0, 0, M, M)
	canvas	:= image.NewGray(r)
	complete = make(chan int)
	RecDraw(canvas, c_left_top, c_right_bottom, depth)
	for i := 0; i < (1 << depth); i++ {
		<- complete
		fmt.Println("completed")
	}

	img := canvas.SubImage(r)
	ShowImage(&img)
}

func RecDraw(canvas *image.Gray, c1, c2 complex128, depth uint) {
	if (depth == 0) {
		go DrawMandelbrot(canvas, c1, c2)
	} else {
		im := (imag(c1) + imag(c2)) / 2
		RecDraw(canvas, c1, complex(real(c2), im), depth - 1)
		RecDraw(canvas, complex(real(c1), im), c2, depth - 1)
	}
}

func DrawMandelbrot(canvas *image.Gray, c1, c2 complex128) {
	fmt.Println("DrawMandelbrot: ", c1, c2)
	for re := real(c1); re < real(c2); re += step {
		for im := imag(c1); im > imag(c2); im -= step {
			c := complex(re, im)
			if InSet(c) {
				p := TranslateToPoint(c)
				canvas.Pix[p.X + p.Y * canvas.Stride] = 255
			}
		}
	}
	complete <- 1
}

func InSet(c complex128) bool {
	z := c
	for i := 0; i < NMAX; i++ {
		z = z * z + c
		if (cmplx.Abs(z) > 2) {
			return false
		}
	}
	return true
}

func ShowImage(img *image.Image) {
	file, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Printf("Could not create file %s", file.Name())
	}
	writer := bufio.NewWriter(file)
	png.Encode(writer, *img)
	writer.Flush()
	file.Close()
}

func TranslateToPoint(c complex128) image.Point {
	var p image.Point
	p.X = (int)((real(c) + 2) * (float64)(M) / rlen + 0.5)
	p.Y = (int)((imag(c) * -1 + 2) * (float64)(M) / rlen + 0.5)
	return p
}