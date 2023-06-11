package main

import "fmt"

func divide(num1, num2 int) (int, int){

quotient := num1 / num2
remainder := num1 % num2

return quotient, remainder

}


func main(){

var num1, num2 int 

fmt.Println("Enter the first number")
fmt.Scanln(&num1)

fmt.Println("Entee the second number")
fmt.Scanln(&num2)

quotient, remainder := divide(num1, num2)

fmt.Printf("The quotient of %d divided by %d is %d\n", num1, num2, quotient)
fmt.Printf("The remainder of %d divided by %d is %d\n",num1, num2, remainder)






}