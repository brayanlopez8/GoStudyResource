package main

import "fmt"

//Calculate function sum 2 value
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
func main() {
	fmt.Println("Go Testing Tutorial")
	result := Calculate(2)
	fmt.Println(result)
}
