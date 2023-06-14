package main

import "fmt"

func findDuplicates(numbers []int) map[int]int {
	//findDuplicates name of funtion defined by us
   //parameter list, which specifies that the function expects a slice of integers as input (numbers)
   //map will contain the duplicates found in the input slice along with their frequencies
	frequency := make(map[int]int)

	// Iterate over the numbers
	for _, num := range numbers {

		frequency[num]++
	}

	duplicates := make(map[int]int)


	for num, count := range frequency {
	
		if count > 1 {
	
			duplicates[num] = count
		}
	}

	return duplicates
}

func main() {
	
	numbers := []int{1, 2, 3, 4, 3, 5, 6, 2, 4}

	duplicates := findDuplicates(numbers)

	
	fmt.Println("Duplicates:")
	for num, count := range duplicates {
		fmt.Printf("%d (count: %d)\n", num, count)
	}
}
