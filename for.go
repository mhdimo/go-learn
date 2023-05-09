package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	for i := 10; i <= 25; i = i + 5 {
		fmt.Println(i)
	}
	fmt.Println("\n")
	for num := 1; num <= 5; num++ {
		if num == 4 {
			continue //continue next iteration // Skips the number when it reach 4
		}
		fmt.Println(num)
	}
	fmt.Println("\n")
	for num := 1; num <= 5; num++ {
		if num == 4 {
			break //exit for loop from here when num reaches 4
		}
		fmt.Println(num)
	}

	fmt.Println("\n")

	//for in range

	colors := [] string{"Red", "Yellow", "Green"}

	for num, val := range colors {
		fmt.Println(num,"-",val)
	}


	// outer loop
    for i :=0; i < 2; i++ {
		// inner loop 
		for j :=0; j < 3; j++ {
			fmt.Println ("i=",i,"j=", j )
		}
  
	}  

}
