package main

import "fmt"

func substract(num1, num2 int)int{

return(num1 - num2)

}

func main(){

var num1, num2 int

fmt.Println("Enter the first number")
fmt.Scanln(&num1)

fmt.Println("Enter the second number")
fmt.Scanln(&num2)

sum := substract(num1, num2)

fmt.Printf("The difference in %d and %d is %d\n" , num1, num2, sum)

}

