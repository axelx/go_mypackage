package main

import "fmt"

func main() {
	a := Add(2, 3)
	fmt.Println(a)
}
func Add(a, b int) int {
	return a + b
}
