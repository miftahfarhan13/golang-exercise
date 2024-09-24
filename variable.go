package main

import "fmt"

func variable() {
	// Ini Komentar
	fmt.Println("Hello World")

	// Variables
	// var
	// -> bisa digunakan di dalam dan di luar fungsi
	// -> deklarasi variable dan penetapan value dapat dilakukan terpisah
	var name string = "Farhan"
	var nickName = "Miftah"

	// :=
	// -> hanya bisa di gunakan di dalam fungsi
	// -> deklarasi variable dan penetapan value tidak dapat dipisah
	age := 29

	// implementasi multiple variable
	var a, b, c = 1, "test", 3

	// implementasi variable block
	var (
		firstName string = "Miftah"
		lastName         = "Farhan"
		weight    int    = 82
	)

	// constanta di golang
	const AGE int = 20

	// constanta di golang dengan variable block
	const (
		X int = 10
		Y     = 3
	)

	fmt.Println(name)
	fmt.Println(nickName)
	fmt.Println(age)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(weight)
}
