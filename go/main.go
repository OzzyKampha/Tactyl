package main

import (
	"fmt"
	"tactyl/tmffi"
)

func main() {
	// Example bit vector for Tsetlin Machine input
	bitVector := []uint8{1, 0, 1, 1, 0, 1}
	result := tmffi.Predict(bitVector)
	fmt.Printf("Input bit vector: %v\n", bitVector)
	fmt.Printf("Number of 1s (dummy TM output): %d\n", result)
}
