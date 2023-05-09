package main

import "fmt"

var num int = 100

func main() {
	// the most common method of declaring a variable in go is by using the keyword `var`
	// var variablename datatype = value;

	var num int = 20           //int variable
	var amount float32 = 54.23 // float variable
	var isvalid bool = true    // bool variable
	var name string = "ciao"   // string variable

	fmt.Println(num, amount, isvalid, name) // output: 20 49.99 true Yash

	// the data type can be omitted whenever declaring a variable
	// so that would be
	// var variablename = value;
	// variables can also be declared without assigning a initial value (which is default)
	// so that would be
	
	// var variablename int
	// variablename = value

	// multible variables can be declared and initialized in the same line
	// so that would be

	// var name1,name2 int = 20,30
	// var amount, name = 49.99, "Yash"

	/*
		using the shorthand syntax

		you can use := to declare and assign a value to a variable without using the keyword var

		so that would be

		num := 20
		amount := 43.22
		isvalid := true
		name := "yash"

		as i explained earlier, we can also declare multiple variables with :=
		num1,num2 := 20,30 // int variables
		amount, name := 32.34, "Yash"
	*/

	fmt.Println(num) // output: 100

	var greet string = "hello"
	fmt.Println(greet) //output: hello

	test()


	//in go we have also constant variables
	//how do we declare it?
	//simply doing
	// const variable = value
	// var name const datatype = value
}


func test() {
	fmt.Println(num) // this is correct because it's a gloabl variable
	// fmt.Println(greet) // error
}