package main

import "fmt"

func isPalindrome(s string) bool {
    
    runes := []rune(s)
    length := len(runes)

    
    for i := 0; i < length/2; i++ {
        if runes[i] != runes[length-1-i] {
            return false
        }
    }

    return true
}

func main() {
   
    strings := []string{"level", "hello", "racecar", "world"}

    for _, str := range strings {
        if isPalindrome(str) {
            fmt.Printf("%s is a palindrome.\n", str)
        } else {
            fmt.Printf("%s is not a palindrome.\n", str)
        }
    }
}
