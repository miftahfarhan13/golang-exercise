package main

import "fmt"

func dataType() {
	var isMarried bool = true
	// int bisa minus
	var myAge int = -28
	// uint tidak bisa minus
	var myHeight uint = 190

	var myWeight float32 = 3.14
	var myName string = "Farhan"

	fmt.Println(isMarried)
	fmt.Println(myAge)
	fmt.Println(myWeight)
	fmt.Println(myName)
	fmt.Println(myHeight)

}
