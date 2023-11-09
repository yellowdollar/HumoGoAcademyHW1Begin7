package main

import "fmt"

func main() {
	var r int

	fmt.Println("Input Radius: ")
	fmt.Scan(&r)

	fmt.Println("L = ", 2*3.14*float32(r))
	fmt.Println("S = ", 3.14*float32(r)*float32(r))
}
