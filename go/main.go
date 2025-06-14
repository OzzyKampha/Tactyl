package main

import (
	"fmt"
	"tactyl/tmffi"
)

func main() {
	numFeatures := 2
	model := tmffi.NewModel(numFeatures)
	defer model.Free()

	// XOR truth table
	inputs := []uint64{0b00, 0b01, 0b10, 0b11}
	targets := []int{0, 1, 1, 0}

	// Train for several epochs
	epochs := 20
	for epoch := 0; epoch < epochs; epoch++ {
		for i, input := range inputs {
			model.Train(input, int32(targets[i]))
		}
	}

	// Test after training
	fmt.Println("\nXOR predictions after training:")
	for i, input := range inputs {
		prediction := model.Predict(input)
		fmt.Printf("Input %02b: predicted=%d, target=%d\n", input, prediction, targets[i])
	}
}
