package main

import "fmt"

func main() {
	x := 5
	y := 4

	// add x and y variables
	resultadd := x + y
	fmt.Println(resultadd) // output 9

	//subtract y from x
	resultsub := x - y
	fmt.Println(resultsub) // output 1

	//modulus of x and y
	resultmod := x % y
	fmt.Println(resultmod) // output 1

	//COMPARISON OPERATORS:
	/*
		Operator 	Description 	Syntax
		== 			Equal 			x == y
		!= 			Not Equal 		x != y
		< 			Less than 		x < y
		<= 			Less than or equal 	x <= y
		> 			Greater than 	x > y
		>= 			Greater than or equal 	x >= y
	*/

	fmt.Println(x == y) // output: false
	fmt.Println(x != y) // output: true
	fmt.Println(x < y)  // output: false
	fmt.Println(x <= y) // output: false
	fmt.Println(x > y)  // output: true
	fmt.Println(x >= y) // output: true


	// LOGICAL OPEARTORS

	/*
		Operator 	Description 														Syntax
		&& 			Logical AND. Returns true if both expressions are true 				expression1 && expression2
		|| 			Logical OR. Returns true if any one of the expressions is true. 	exp1 || exp2
		! 			Logical NOT. Returns true if an expression is false. Returns false if the expression is true. 	! exp 
	*/



	// BITWISE OPERATORS

	/*
		Operator	Description
		&			bitwise AND
		|			bitwise OR	
		^			bitwise XOR 
		<<			left shift
		>>			right shift
	*/

	var t,u = 3,5
	var z int

	//Bitwise AND
	z = t & u
	fmt.Println("Bitwise AND ", z) // output: 1
	
    //Bitwise OR
	z = t | u
	fmt.Println("Bitwise OR ", z) //output: 7
	
    //Bitwise XOR
	z = t ^ u
	fmt.Println("Bitwise XOR ", z) //output: 6


	//Assignment Operators
	/*
	Operator 	Description 				Syntax
	= 			Assign value 				x = y
	+= 			Add and assign 				x += y same as x = x + y
	-= 			Subtract and assign 		x -= y same as x = x - y
	*= 			Multiply and assign 		x *= y same as x = x * y
	/= 			Divide and assign 			x /= y same as x = x/y
	%= 			Divide and assign modulus 	x %= y same as x = x%y 
	*/

	//Assign
	x = y
	fmt.Println(" = ", x) //output: 20

	// Add and assign
	x = 15
	x += y
	fmt.Println("+= ", x) //output: 35

	// Subtract and assign
	x = 25
	x -= y
	fmt.Println("-=", x) //output: 5

	//Multiply and assign
	x = 2
	x *= y
	fmt.Println("*=", x) //output: 40

	//Divide and assign
	x = 100
	x /= y
	fmt.Println("/=", x) // output: 5

	//Divide and assign modulus
	x = 30
	x %= y
	fmt.Println(" %= ", x) //output: 10
}
