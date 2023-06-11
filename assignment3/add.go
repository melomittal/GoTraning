package main

import "fmt"

func add(num1, num2 int)int {

return(num1 + num2)

}

func main(){

var num1, num2 int 

fmt.Println("Enter the first number")
fmt.Scanln(&num1)

fmt.Println("Enter the second number")
fmt.Scanln(&num2)

sum := add(num1, num2)


fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

}