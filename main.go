package main

import (
	"fmt"
	"strconv"
)

var state bool
var operation string
var input1 string
var input2 string

func main() {

	state = true
	for {
		calculator()
		if !state {
			break
		}
	}
}
func calculator() {
	for {
		fmt.Print("Enter your operation (a)dd, (s)ubtract, (m)ultiply, (d)ivide: ")
		fmt.Scanln(&operation)
		if operation == "a" || operation == "s" || operation == "m" || operation == "d" {
			break
		} else {
			fmt.Println("Invalid operation choice. Try again!")
		}
	}
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&input1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&input2)
	num1, _ := strconv.ParseFloat(input1, 32)
	num2, _ := strconv.ParseFloat(input2, 32)
	switch operation {
	case "a":
		fmt.Println(num1 + num2)
	case "s":
		fmt.Println(num1 - num2)
	case "m":
		fmt.Println(num1 * num2)
	case "d":
		if num2 == 0 {
			fmt.Println("Divide by Zero Error")
		} else {
			fmt.Println(num1 / num2)
		}
	}

	var again string
	fmt.Print("Go Again? y/n: ")
	fmt.Scanln(&again)
	if again == "y" {
		state = true
	} else {
		state = false
		fmt.Println("Bye!")
	}
}
