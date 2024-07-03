package main

import (
	"fmt"
)

func main() {
	// user input with single name without whitespace
	fmt.Println("Hey, What's Your Name")

	var name string
	fmt.Scan(&name)
	fmt.Println("Hello Mr:", name)

}
