package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Arie() float64
	Perimetru() float64
}

type Cerc struct {
	raza float64
}

type Dreptunghi struct {
	latime, inaltime float64
}

type Triunghi struct {
	a, b, c float64
}

func (c Cerc) Arie() float64 {
	return math.Pi * c.raza * c.raza
}

func (c Cerc) Perimetru() float64 {
	return 2 * math.Pi * c.raza
}

func (d Dreptunghi) Arie() float64 {
	return d.latime * d.inaltime
}

func (d Dreptunghi) Perimetru() float64 {
	return 2 * (d.latime + d.inaltime)
}

func (t Triunghi) Arie() float64 {
	p := (t.a + t.b + t.c) / 2.0
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (t Triunghi) Perimetru() float64 {
	return t.a + t.b + t.c
}

func main() {
	forme := []Forma{
		Cerc{raza: 4},
		Dreptunghi{latime: 5, inaltime: 3},
		Triunghi{a: 3, b: 4, c: 5},
	}

	var ariaTotala float64
	var maxAria float64 = -1
	var formaMaxima Forma

	for _, f := range forme {
		a := f.Arie()
		p := f.Perimetru()
		ariaTotala += a

		fmt.Printf("Tip: %-12T | Arie: %6.2f | Perimetru: %6.2f\n", f, a, p)

		if a > maxAria {
			maxAria = a
			formaMaxima = f
		}
	}

	fmt.Printf("Aria totala: %.2f\n", ariaTotala)
	fmt.Printf("aria maxima %T (cu aria de %.2f)\n", formaMaxima, maxAria)
}
