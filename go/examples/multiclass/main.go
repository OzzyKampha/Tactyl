package main

import (
	"fmt"
	"tactyl/tmffi"
)

func main() {
	// Create a Tsetlin Machine for a 3-class problem
	tm := tmffi.NewModel(3, 3, 3.9, 100, 50, 20, 0.0)
	if tm == nil {
		panic("Failed to create Tsetlin Machine")
	}
	defer tm.Free()

	// Training data: 3 classes with distinct patterns
	trainingData := []struct {
		input  []bool
		target int
	}{
		{[]bool{true, false, false}, 0}, // Class 0
		{[]bool{false, true, false}, 1}, // Class 1
		{[]bool{false, false, true}, 2}, // Class 2
	}

	fmt.Println("Training the Tsetlin Machine...")
	fmt.Println("Initial predictions:")
	for _, data := range trainingData {
		pred := tm.Predict(data.input)
		fmt.Printf("  Input: %v, Target: %d, Predicted: %d\n", data.input, data.target, pred)
	}

	// Train for 10 epochs
	for epoch := 0; epoch < 10; epoch++ {
		for _, data := range trainingData {
			tm.Train(data.input, data.target)
		}

		if epoch%2 == 0 {
			fmt.Printf("\nAfter epoch %d:\n", epoch+1)
			for _, data := range trainingData {
				pred := tm.Predict(data.input)
				fmt.Printf("  Input: %v, Target: %d, Predicted: %d\n", data.input, data.target, pred)
			}
		}
	}

	// Test with some new patterns
	fmt.Println("\nTesting with new patterns:")
	testPatterns := [][]bool{
		{true, true, false}, // Should be closer to class 0
		{false, true, true}, // Should be closer to class 1
		{true, false, true}, // Should be closer to class 2
	}

	for _, pattern := range testPatterns {
		pred := tm.Predict(pattern)
		fmt.Printf("  Input: %v, Predicted class: %d\n", pattern, pred)
	}

	// Print debug information
	fmt.Println("\nDebug information:")
	tm.PrintStateHistogram()

	// Calculate F1 accuracy
	fmt.Println("\nF1 Accuracy:")
	truePositives := make([]int, 3)
	falsePositives := make([]int, 3)
	falseNegatives := make([]int, 3)

	for _, data := range trainingData {
		pred := tm.Predict(data.input)
		if pred == data.target {
			truePositives[data.target]++
		} else {
			falsePositives[pred]++
			falseNegatives[data.target]++
		}
	}

	for i := 0; i < 3; i++ {
		precision := float64(truePositives[i]) / float64(truePositives[i]+falsePositives[i])
		recall := float64(truePositives[i]) / float64(truePositives[i]+falseNegatives[i])
		f1 := 2 * precision * recall / (precision + recall)
		fmt.Printf("  Class %d: Precision = %.2f, Recall = %.2f, F1 = %.2f\n", i, precision, recall, f1)
	}
	tm.PrintStateHistogram()
}
