package main

import "fmt"

func main() {
	fmt.Println("We are learning array in go lang")

	//Syntax of array this example
	var name [5]string
	name[0] = "Shaikh"
	name[1] = "Sufiyan"
	name[2] = "xyz"

	fmt.Println("Name of Person is:", name)

	//Find the Array index element access this Syntax
	fmt.Println("value of name at 1 index is:", name[1])

	//Array Initialization Syntax
	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is:", number)

	//Find the Length of Array this Syntax
	fmt.Println("Length of the Array is:", len(number))
}
