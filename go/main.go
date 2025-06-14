package main

import (
	"fmt"
	"tactyl/tmffi"
)

func main() {
	numFeatures := 2
	numClauses := 100
	t := 50
	s := 3.0
	model := tmffi.NewModel(numFeatures, numClauses, t, s)
	defer model.Free()

	// AND truth table
	inputs := []uint64{0b00, 0b01, 0b10, 0b11}
	targets := []int{0, 0, 0, 1}

	// Train for several epochs
	epochs := 100
	for epoch := 0; epoch < epochs; epoch++ {
		for i, input := range inputs {
			model.Train(input, int32(targets[i]))
		}
	}

	// Print automata state distribution
	model.PrintStateHistogram()

	// Test after training with debug info
	fmt.Println("\nAND predictions and clause firing after training:")
	for i, input := range inputs {
		prediction := model.Predict(input)
		fmt.Printf("Input %02b: predicted=%d, target=%d\n", input, prediction, targets[i])
	}
}
