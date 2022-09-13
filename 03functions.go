package main

import "fmt"

func main() {

	welcome()

	addition(5, 3)

	result := multiply(4, 5)
	fmt.Println(result)

	sum, multiple := maths(5, 5)
	fmt.Println(sum, "and", multiple)

	sumNew, multipleNew := mathsNew(6, 6)
	fmt.Println(sumNew, "and", multipleNew)

}

func mathsNew(i1, i2 int) (s int, m int) {
	s = i1 + i2
	m = i1 * i2
	return s, m
}

func maths(i1, i2 int) (int, int) {
	return i1 + i2, i1 * i2
}

func multiply(i1, i2 int) int {
	return i1 * i2
}

func addition(i1 int, i2 int) {

	fmt.Println(i1 + i2)

}

func welcome() {

	fmt.Println("Welcome to functions in Go Lang")

}
