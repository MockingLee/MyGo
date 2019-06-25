package main

import "fmt"

func first() {
	defer fmt.Println("first")
	fmt.Println("first func run")
}
func second() {
	defer fmt.Println("second")
	fmt.Println("second func run")
}

func main() {
	var x int
	var x_ptr *int

	x = 10
	x_ptr = &x

	fmt.Println(&x_ptr)
}