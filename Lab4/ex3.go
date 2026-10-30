package main

import "fmt"

type Persoana struct {
	Nume   string
	Varsta int
}

func (p Persoana) Prezentare() {
	fmt.Println(p.Nume, p.Varsta)
}

type Student struct {
	Persoana
	Grupa int
}

func (s Student) String() {
	fmt.Println(s.Nume, s.Varsta, s.Grupa)
}

func main() {

	s := Student{
		Persoana: Persoana{
			Nume:   "Ana Popescu",
			Varsta: 21,
		},
		Grupa: 342,
	}

	s.Prezentare()
	fmt.Println(s)
}
