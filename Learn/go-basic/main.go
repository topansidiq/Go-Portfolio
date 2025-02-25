package main

import "fmt"

func main() {

	// HELLO WORLD
	fmt.Println("Hello, World!")

	// VARIABLES
	/*
		Keyword var
		Declaration with data type
	*/
	var firstName = "Topan"
	var lastName string = "Sidiq"

	var fullName = firstName + " " + lastName

	fmt.Printf("First name: %s\n", firstName)
	fmt.Printf("Last name: %s\n", lastName)
	fmt.Printf("Full name: %s\n", fullName)

	/*
		Without data type
		:=
	*/
	var x int64 = 901
	y := x - 1

	fmt.Println(y)

	y = 100
	y = 101

	fmt.Println(y)

	/*
		Multi variable
	*/
	var first, second, third string
	first, second, third = "First", "Second", "Third"

	fmt.Printf("%s %s %s\n", first, second, third)

	a, b, c, d := "cm", 45, 90.1, true

	fmt.Println(a, b, c, d)

	/*
		Underscore variable
	*/
	_ = "Black hole"

	name, _ := "Topan Sidiq", 123
	fmt.Println(name)

	/*
		Keyword new
	*/

	name2 := new(string)

	fmt.Println(name2)
	fmt.Println(*name2)

	// DATA TYPES

	/*
		Integer
		uint = positif
		int = positif dan negatif
	*/
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var minUint uint = 0
	var minUint8 uint8 = 0
	var minUint16 uint16 = 0
	var minUint32 uint32 = 0
	var minUint64 uint64 = 0

	fmt.Println(minUint, minUint8, minUint16, minUint32, minUint64)

	var maxUint uint = 18446744073709551615
	var maxUint8 uint8 = 255
	var maxUint16 uint16 = 65535
	var maxUint32 uint32 = 4294967295
	var maxUint64 uint64 = 18446744073709551615

	fmt.Println(maxUint, maxUint8, maxUint16, maxUint32, maxUint64)
}
