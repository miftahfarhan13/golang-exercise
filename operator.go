package main

import "fmt"

func operator() {
	var a, b, c = 10, 11, 12
	fmt.Println(a + b)
	fmt.Println(a - c)
	fmt.Println(b * a)
	fmt.Println(b / c)
	fmt.Println(c % a)

	// increment
	i := 5
	i++

	// decrement
	j := 4
	j--

	// operator assignment
	var d, e, f = 13, 14, 15
	d += 5
	e -= 5
	f *= 6
	d /= 8
	e %= 2

	// operator perbandingan
	var x = 10
	var y = 11

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x >= y)

	// operator logika
	var and = x > 10 && y < 10
	var or = x > 10 || y < 10

	fmt.Print(and)
	fmt.Print(or)

}
