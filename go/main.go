package main

import (
	"fmt"
	"tactyl/tmffi"
)

func printTrainingExample(input []uint8, target int32, prediction int32) {
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Target: %d, Prediction: %d\n", target, prediction)
	fmt.Printf("Correct: %v\n\n", target == prediction)
}

func main() {
	// Configure TM parameters
	tmffi.SetSpecificity(10.0) // s = 10
	tmffi.SetThreshold(250.0)  // threshold = 250 (half of 500 clauses)

	// XOR training data
	trainingData := [][]uint8{
		{0, 0}, // 0
		{0, 1}, // 1
		{1, 0}, // 1
		{1, 1}, // 0
	}
	targets := []int32{0, 1, 1, 0}

	fmt.Println("Starting Tsetlin Machine Training (XOR)...")
	fmt.Println("============================================")

	// Train the TM for multiple epochs
	epochs := 100
	for epoch := 0; epoch < epochs; epoch++ {
		correct := 0
		fmt.Printf("\nEpoch %d:\n", epoch+1)
		fmt.Println("-------------------")

		for i, input := range trainingData {
			// Train
			tmffi.Train(input, targets[i])

			// Predict
			prediction := tmffi.Predict(input)

			// Print results
			printTrainingExample(input, targets[i], prediction)

			if prediction == targets[i] {
				correct++
			}
		}

		accuracy := float64(correct) / float64(len(trainingData)) * 100
		fmt.Printf("Epoch %d Accuracy: %.2f%%\n", epoch+1, accuracy)
	}

	// Test with XOR examples
	fmt.Println("\nTesting with XOR Examples:")
	fmt.Println("=========================")

	testData := [][]uint8{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	testTargets := []int32{0, 1, 1, 0}

	for i, input := range testData {
		prediction := tmffi.Predict(input)
		fmt.Printf("\nTest Example %d:\n", i+1)
		printTrainingExample(input, testTargets[i], prediction)
	}
}
