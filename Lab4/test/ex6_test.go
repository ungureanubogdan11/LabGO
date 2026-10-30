package main

import (
	"errors"
	"testing"
)

func TestDepuneValid(t *testing.T) {
	c := &ContBancar{Sold: 100.0}
	err := c.Depune(50.0)
	if err != nil || c.Sold != 150.0 {
		t.Fail()
	}
}

func TestDepuneInvalida(t *testing.T) {
	c := &ContBancar{Sold: 100.0}
	err := c.Depune(-10.0)
	if !errors.Is(err, ErrSumaInvalida) || c.Sold != 100.0 {
		t.Fail()
	}
}

func TestRetrageValid(t *testing.T) {
	c := &ContBancar{Sold: 100.0}
	err := c.Retrage(40.0)
	if err != nil || c.Sold != 60.0 {
		t.Fail()
	}
}

func TestRetrageInsuficient(t *testing.T) {
	c := &ContBancar{Sold: 100.0}
	err := c.Retrage(150.0)
	if !errors.Is(err, ErrSoldInsuficient) || c.Sold != 100.0 {
		t.Fail()
	}
}

func TestRetrageInvalida(t *testing.T) {
	c := &ContBancar{Sold: 100.0}
	err := c.Retrage(-20.0)
	if !errors.Is(err, ErrSumaInvalida) || c.Sold != 100.0 {
		t.Fail()
	}
}
