package main

import (
	"fmt"
)

// Struct Person dengan method
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Metode untuk struct Person
func (p Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) IncreaseAge() {
	p.Age++
}

// Struct Rectangle dengan method
type Rectangle struct {
	Length float64
	Width  float64
}

// Metode untuk struct Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func main() {
	// Menggunakan struct Person
	person := Person{
		FirstName: "Kelvin",
		LastName:  "Adiyasa",
		Age:       19,
	}

	fmt.Println("Nama Lengkap:", person.GetFullName())
	fmt.Println("Usia Awal:", person.Age)
	person.IncreaseAge()
	fmt.Println("Usia setelah IncreaseAge:", person.Age)

	// Menggunakan struct Rectangle
	rect := Rectangle{
		Length: 10,
		Width:  5,
	}

	fmt.Println("Luas Rectangle:", rect.Area())
}
