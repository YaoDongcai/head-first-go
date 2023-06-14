package main

import "fmt"

// fun main
func main() {
	var count int = 20
	unitWeight := 0.4
	totalWeight := float64(count) * unitWeight
	fmt.Println(count, "cans weigh", totalWeight, "kilograms")
}
