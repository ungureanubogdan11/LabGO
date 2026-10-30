package main

import (
	"errors"
	"fmt"
)

var (
	ErrSumaInvalida    = errors.New("suma invalida")
	ErrSoldInsuficient = errors.New("sold insuficient")
)

type ContBancar struct {
	Titular string
	Sold    float64
}

func (c *ContBancar) Depune(suma float64) error {
	if suma <= 0 {
		return ErrSumaInvalida
	}
	c.Sold += suma
	return nil
}

func (c *ContBancar) Retrage(suma float64) error {
	if suma <= 0 {
		return ErrSumaInvalida
	}
	if suma > c.Sold {
		return ErrSoldInsuficient
	}
	c.Sold -= suma
	return nil
}

func main() {
	cont := &ContBancar{
		Titular: "alex",
		Sold:    500.0,
	}

	fmt.Println(cont.Titular, cont.Sold)

	operatii := []struct {
		tip  string
		suma float64
	}{
		{"depunere", 200.0},
		{"depunere", -50.0},
		{"retragere", 150.0},
		{"retragere", 1000.0},
		{"retragere", -20.0},
	}

	for _, op := range operatii {
		var err error

		fmt.Print(op.tip + " ")
		if op.tip == "depunere" {
			fmt.Println(op.suma)
			err = cont.Depune(op.suma)
		} else {
			fmt.Println(op.suma)
			err = cont.Retrage(op.suma)
		}

		if err != nil {
			if errors.Is(err, ErrSumaInvalida) {
				fmt.Println(err)
			} else if errors.Is(err, ErrSoldInsuficient) {
				fmt.Println(err)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(cont.Sold)
		}
	}
}
