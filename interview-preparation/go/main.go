package main

import (
	"fmt"
	"interview-preparation/go/src"
)

func main() {
	fmt.Println("🚀 Go Interview Preparation Project Initialized!")
	
	a, b := 10, 5
	fmt.Printf("Sum of %d and %d is: %d\n", a, b, src.Add(a, b))
	
	n := 5
	fmt.Printf("Factorial of %d is: %d\n", n, src.Factorial(n))
}
