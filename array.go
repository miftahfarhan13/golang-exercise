package main

import "fmt"

func array() {
	var a1 = [3]int{}                //tidak di inisialisasi
	a2 := [4]int{400, 500}           //di inisialisasi sebagian
	a3 := [4]int{400, 500, 600, 700} //di inisialisasi semuanya

	var a4 = [...]int{800, 900}
	var a5 = [...]string{"go", "javascript", "node"}

	// tipe data slice
	a6 := []string{"string", "string2"}

	// akses index array
	fmt.Println(a5[0])

	// akses panjang array
	fmt.Println(len(a1))
	// akses kapasitas slice
	fmt.Println(cap(a6))

	// membuat slice dari array
	// index low sampe sebelum index high
	slice := a3[1:3]
	// index low sampe index terakhir
	slice2 := a3[1:]
	// index pertama sampe sebelum index high
	slice3 := a3[:3]
	// index pertama sampe index terakhir
	slice4 := a3[:]

	// pass by reference
	slice[0] = 100

	// membuat slice dengan make
	sliceMake := make([]int, 5)
	sliceMake[0] = 1
	sliceMake[1] = 2

	// copy slice
	copySlice := make([]int, len(sliceMake), cap(sliceMake))
	copy(copySlice, sliceMake)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a6)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(sliceMake)

}
