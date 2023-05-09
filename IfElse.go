package main

import "fmt"

func main() {

	x := 8
	y := 10

	compare(x, y)

	var temp float32 = 23.67

	checktemp(temp)

	compare2(x,y)
}

func compare(a int, b int) {
	if a < b {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than y")
	}
}

func checktemp(temp float32) {
	if temp < 15 {
		fmt.Println("cold")
	} else if temp < 30 {
		fmt.Println("Moderate")
	} else {
		fmt.Println("Hot")
	}
}

func compare2(a int, b int) {
	
	if (a != b) {
        if (a < b) {
            fmt.Println("x is less than y");
        } else if (a > b) {
            fmt.Println("x is greater than y");
        }        
    } else {
            fmt.Println("x is equal to y");
        }
}