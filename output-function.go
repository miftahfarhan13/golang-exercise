package main

import "fmt"

func OutputFunction() {
	var firstName, lastName = "Miftah", "Farhan"
	// print tanpa enter, tanpa spasi
	fmt.Print("")
	// print dengan enter
	fmt.Print("", "\n")
	// print dengan spasi
	fmt.Print(firstName, " ", lastName)

	// printLn dengan menambahkan spasi antar argumen dan baris baru di akhir
	fmt.Println(firstName, lastName)

	name := "Farhan"
	age := 20
	// print yang dapat mengeluarkan tipe data dari sebuat variable
	fmt.Printf("value name adalah : %v dan tipe data name adalah %T\n", name, name)
	fmt.Printf("value age adalah : %v dan tipe data age adalah %T\n", age, age)

	// %v mencetak value dalam format default
	// %#v mencetak value dalam format Go-syntax
	// %T mencetak tipe data value
	// %% Mencetak tanga %
}
