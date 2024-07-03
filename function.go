package main

import "fmt"

func SimpleFunction() {
	fmt.Println("simplefunction")
}
func add(a, b int) (result int) {
	result = a + b
	return
}
func multi(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	SimpleFunction()

	ans1 := add(3, 4)
	fmt.Println("Addition of two no is:", ans1)

	ans2 := multi(5, 4)
	fmt.Println("Multiply by two no is:", ans2)
}
