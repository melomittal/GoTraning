package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// numbers is defined and initialized with a sequence of numbers.
	even := make([]int, 0)
	odd := make([]int, 0)
	// even and odd are two empty slices. The make funtion is being used to initiate a slice.

	for _, num := range numbers {
		// for(loop)_(to discard the index)num(current element) represents iteration over in a loop
		if num%2 == 0 {
			// is a condition that checks if num is even or /2. If condition true, number is even. if false then its odd.
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
		// append function is used to add even and odd numbers to their respective slices (even and odd slices).
	}

	fmt.Println("Even numbers:", even)
	fmt.Println("Odd numbers:", odd)
}




