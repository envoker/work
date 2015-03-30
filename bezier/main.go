package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func main() {

	if err := start(); err != nil {
		fmt.Println(err.Error())
	}
}

func start() error {

	rect := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{512, 512},
	}

	var m *image.RGBA

	m = image.NewRGBA(rect)

	//imageFill(m, color.RGBA{255, 255, 0, 255})
	imageFill(m, color.White)

	r := newRand()
	for i := 0; i < 5; i++ {
		makeCurve(m, r)
	}

	err := imageWriteToPNG(m, "./test.png")
	if err != nil {
		return err
	}

	return nil
}

func imageFill(m *image.RGBA, c color.Color) {

	c1 := color.RGBAModel.Convert(c).(color.RGBA)

	dst := m.Pix

	src := []byte{
		c1.R,
		c1.G,
		c1.B,
		c1.A,
	}

	for len(dst) > 0 {
		copy(dst, src)
		dst = dst[4:]
	}
}

func imageWriteToPNG(m image.Image, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	err = png.Encode(w, m)
	if err != nil {
		return err
	}

	return nil
}

func makeCurve(m *image.RGBA, r *rand.Rand) error {

	bounds := m.Bounds()

	var (
		minX = int(float64(bounds.Dx()) * 0.1)
		maxX = int(float64(bounds.Dx()) * 0.9)

		minY = int(float64(bounds.Dy()) * 0.1)
		maxY = int(float64(bounds.Dy()) * 0.9)
	)

	var (
		dX = maxX - minX
		dY = maxY - minY
	)

	n := 10
	ps := make([]Point, (n*(n+1))/2)

	for i := 0; i < n; i++ {
		ps[i] = Point{
			X: float64(minX + r.Intn(dX)),
			Y: float64(minY + r.Intn(dY)),
		}
	}

	t := 0.0
	dt := 0.001
	for t < 1.0 {
		draw(m, r, ps, n, t)
		t += dt
	}

	return nil
}

func draw(m *image.RGBA, r *rand.Rand, ps []Point, n int, t float64) error {

	if n > 1 {

		a := 8.0 * t * (1 - t)

		for i := 0; i < n-1; i++ {

			ps[n+i] = PointLerp(ps[i], ps[i+1], t)

			ps[n+i].X += a * (1.0 - r.Float64()*2.0)
			ps[n+i].Y += a * (1.0 - r.Float64()*2.0)
		}
		return draw(m, r, ps[n:], n-1, t)
	}

	x := int(Round(ps[0].X))
	y := int(Round(ps[0].Y))

	m.Set(x, y, color.Black)

	return nil
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
