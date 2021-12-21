package main

import "fmt"

func main() {
	// add two integers
	fmt.Println(add(1,2))
	// add two floats
	fmt.Println(add(1.0,2.3))
	// add two strings
	fmt.Println(add("1.0","2.3"))
}
func add[T int|int8|int16|int32|int64|float32|float64|string](a,b T) T {
	return a+b
}
