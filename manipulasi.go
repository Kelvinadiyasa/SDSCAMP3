package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) UpdateName(newFirstName, newLastName string) {
	p.FirstName = newFirstName
	p.LastName = newLastName
}

func (p *Person) IncreaseAge() {
	p.Age++
}

func main() {
	// Manipulasi data menggunakan nilai (tidak menggunakan pointer)
	person := Person{
		FirstName: "Kelvin",
		LastName:  "Adiyasa",
		Age:       19,
	}

	fmt.Println("Data Orang Sebelum Manipulasi:")
	fmt.Println("Nama Lengkap:", person.FirstName, person.LastName)
	fmt.Println("Usia:", person.Age)

	person.UpdateName("Taylor", "Swift") // Mengubah nama menggunakan nilai
	person.IncreaseAge()                 // Menambah usia menggunakan pointer

	fmt.Println("\nData Orang Setelah Manipulasi (menggunakan nilai):")
	fmt.Println("Nama Lengkap:", person.FirstName, person.LastName)
	fmt.Println("Usia:", person.Age)

	// Manipulasi data menggunakan pointer
	personPtr := &Person{
		FirstName: "Cristiano",
		LastName:  "Ronaldo",
		Age:       25,
	}

	fmt.Println("\nData Orang Sebelum Manipulasi (menggunakan pointer):")
	fmt.Println("Nama Lengkap:", personPtr.FirstName, personPtr.LastName)
	fmt.Println("Usia:", personPtr.Age)

	personPtr.UpdateName("Mama", "Mia") // Mengubah nama menggunakan nilai
	personPtr.IncreaseAge()             // Menambah usia menggunakan pointer

	fmt.Println("\nData Orang Setelah Manipulasi (menggunakan pointer):")
	fmt.Println("Nama Lengkap:", personPtr.FirstName, personPtr.LastName)
	fmt.Println("Usia:", personPtr.Age)
}
