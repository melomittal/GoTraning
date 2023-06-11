package main

import "fmt"

func multiplication(num1, num2 int) int {

	return(num1 * num2)

}

func main(){

	var num1, num2 int

	fmt.Println("Enter the first number")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number")
	fmt.Scanln(&num2)

	sum := multiplication(num1, num2)

	fmt.Printf("The multiplication of %d and %d is %d\n", num1, num2, sum)



}