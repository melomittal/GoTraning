package main

import (
	"fmt"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		panic("Cannot divide by zero")
	}
	return a / b
}

func main() {
	var num1, num2 float64

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println("----- Results -----")
	fmt.Printf("Sum: %.2f\n", add(num1, num2))
	fmt.Printf("Difference: %.2f\n", subtract(num1, num2))
	fmt.Printf("Product: %.2f\n", multiply(num1, num2))
	fmt.Printf("Quotient: %.2f\n", divide(num1, num2))
}
