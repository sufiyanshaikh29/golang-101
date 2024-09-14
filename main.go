package main

import "fmt"

func main() {
	fmt.Println("Learn Slice for Go language ")

	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	fmt.Println("number is:", numbers)
	fmt.Printf("numbers has data type is: %T\n", numbers)
	fmt.Println("length of Numbers is:", len(numbers))

}
