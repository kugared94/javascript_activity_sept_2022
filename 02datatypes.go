package main

import "fmt"

func main() {

	fmt.Println("Datatypes")
	// declare a variable in go
	// var variableName datatype
	var score int
	var cityName string
	var check bool
	var lifeline = 88.88

	score = 100
	cityName = "Kajang"
	check = true

	fmt.Println(score)
	fmt.Println(cityName)
	fmt.Println(check)
	fmt.Println(lifeline)

	fmt.Println("Score:", score)
	fmt.Println("City Name:", cityName)
	fmt.Println("Check:", check)
	fmt.Println("Lifeline:", lifeline)
	fmt.Println("My score is", score, "and i live in", cityName)

	// use format specifier
	fmt.Printf("My score is %v. I live in %v. \n", score, cityName)
	fmt.Printf("My lifeline is %v. \n", lifeline)

	// %T format specifier -> get datatype of variable
	fmt.Printf("Datatype of lifeline variable is %V \n", lifeline)

	// auto type inference
	// let go detect the datatype of the variable automatically
	// use :=

	autoScore := "PASS"

	fmt.Printf("My autoScore is %v. \n", autoScore)
	fmt.Printf("Datatype of lifeline variable is %V \n", lifeline)

}
