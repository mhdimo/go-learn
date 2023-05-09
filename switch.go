package main

import "fmt"

func main() {
	var month int = 5

	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28 days")
	}

	thismonth := "May"

	switch thismonth {
	case "January":
		fmt.Println("1st month")
	case "February":
		fmt.Println("2nd month")
	case "March":
		fmt.Println("3rd month")
	case "April":
		fmt.Println("4th month")
	case "May":
		fmt.Println("5th month")
	case "June":
		fmt.Println("6th month")
	case "July":
		fmt.Println("7th month")
	case "August":
		fmt.Println("8th month")
	case "September":
		fmt.Println("9th month")
	case "October":
		fmt.Println("10th month")
	case "November":
		fmt.Println("11th month")
	case "December":
		fmt.Println("12th month")
	}

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Sunday")
	}
}
