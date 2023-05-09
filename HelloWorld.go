package main

import "fmt" //import library

func main() {
	fmt.Println("Hello World") // PrintLn is the equivalent of printf in c++
}


/*
How to run a go program?
	$ go run HelloWorld.go
	this will normally execute the code provided to the compiler


What about if we want binary files then? We will need to build this program
	$ go build HelloWorld.go
	$ ls
	HelloWorld HelloWorld.go

	$ ./HelloWorld
	Hello World
*/