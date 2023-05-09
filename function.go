package main

import "fmt"

func greet() {
	fmt.Println("Hello")
}

func sum(a int, b int) int {
	return a + b
}

func sumsub(a int, b int) (int, int) {
	sum := a + b
	sub := a - b

	return sum, sub
}

func GetDefaults() (int, string) {
	return 0, "unknown"
}

func main() {
	greet()

	result := sum(10, 10)
	fmt.Println(result)

	sum, subtract := sumsub(38943, 3984)

	fmt.Println(sum, subtract)

	value, name := GetDefaults()

	fmt.Println(value, name)
}
