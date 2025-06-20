package main

import (
	"fmt"
	"strings"
	"tactyl/tmffi"
)

// Feature bits for shape recognition:
// bit 0: Has straight lines (1) or curved lines (0)
// bit 1: Has corners (1) or no corners (0)
// bit 2: Is symmetric (1) or asymmetric (0)
// Class 0: Circle (curved, no corners, symmetric)
// Class 1: Square (straight, corners, symmetric)
// Class 2: Triangle (straight, corners, symmetric)
// Class 3: Rectangle (straight, corners, symmetric)
// Class 4: Oval (curved, no corners, symmetric)
// Class 5: Star (straight, corners, asymmetric)
// Class 6: Heart (curved, no corners, asymmetric)
// Class 7: Pentagon (straight, corners, symmetric)

func main() {
	numFeatures := 3   // 3 feature bits
	numClauses := 1000 // Reduced from 2000 to avoid overfitting
	numClasses := 8    // 8 different shapes
	t := 500           // Reduced threshold for better learning
	s := 2.0           // Reduced specificity for more balanced learning
	model := tmffi.NewModel(numFeatures, numClasses, 3.9, numClauses, numClauses, uint8(t), float32(s))
	defer model.Free()

	// Training data for shape recognition
	inputs := [][]bool{
		// Circle variations (curved, no corners, symmetric)
		{false, false, true}, // Basic circle
		{false, false, true}, // Circle with slight variation
		{false, false, true}, // Circle with another variation
		{false, false, true}, // Circle with both variations

		// Square variations (straight, corners, symmetric)
		{true, true, true}, // Basic square
		{true, true, true}, // Square with slight variation
		{true, true, true}, // Square with another variation
		{true, true, true}, // Square with different variation

		// Triangle variations (straight, corners, asymmetric)
		{true, true, false}, // Basic triangle
		{true, true, false}, // Triangle with slight variation
		{true, true, false}, // Triangle with another variation
		{true, true, false}, // Triangle with both variations

		// Rectangle variations (straight, corners, symmetric)
		{true, true, true}, // Basic rectangle
		{true, true, true}, // Rectangle with slight variation
		{true, true, true}, // Rectangle with another variation
		{true, true, true}, // Rectangle with different variation

		// Star variations (straight, corners, asymmetric)
		{true, true, false}, // Basic star
		{true, true, false}, // Star with slight variation
		{true, true, false}, // Star with another variation
		{true, true, false}, // Star with different variation

		// Diamond variations (straight, corners, symmetric)
		{true, true, true}, // Basic diamond
		{true, true, true}, // Diamond with slight variation
		{true, true, true}, // Diamond with another variation
		{true, true, true}, // Diamond with different variation

		// Hexagon variations (curved, corners, symmetric)
		{false, true, true}, // Basic hexagon
		{false, true, true}, // Hexagon with slight variation
		{false, true, true}, // Hexagon with another variation
		{false, true, true}, // Hexagon with different variation

		// Octagon variations (curved, corners, symmetric)
		{false, true, true}, // Basic octagon
		{false, true, true}, // Octagon with slight variation
		{false, true, true}, // Octagon with another variation
		{false, true, true}, // Octagon with different variation
	}

	// Target classes for each input (0-7 for 8 shapes)
	targets := []int{
		// Circle targets
		0, 0, 0, 0,
		// Square targets
		1, 1, 1, 1,
		// Triangle targets
		2, 2, 2, 2,
		// Rectangle targets
		3, 3, 3, 3,
		// Star targets
		4, 4, 4, 4,
		// Diamond targets
		5, 5, 5, 5,
		// Hexagon targets
		6, 6, 6, 6,
		// Octagon targets
		7, 7, 7, 7,
	}

	// Train for several epochs
	epochs := 1000 // Reduced epochs to avoid overfitting
	fmt.Printf("Training for %d epochs...\n", epochs)
	bestAccuracy := 0.0
	bestEpoch := 0

	for epoch := 0; epoch < epochs; epoch++ {
		for i, input := range inputs {
			model.Train(input, targets[i])
		}

		// Print accuracy every 50 epochs
		if (epoch+1)%50 == 0 {
			accuracy := calculateAccuracy(model, inputs, targets)
			fmt.Printf("Epoch %d: Training Accuracy = %.2f%%\n", epoch+1, accuracy*100)
			if accuracy > bestAccuracy {
				bestAccuracy = accuracy
				bestEpoch = epoch + 1
			}
		}
	}

	// Print automata state distribution
	model.PrintStateHistogram()

	// Calculate and print final training accuracy
	trainingAccuracy := calculateAccuracy(model, inputs, targets)
	fmt.Printf("\nFinal Training Accuracy: %.2f%%\n", trainingAccuracy*100)
	fmt.Printf("Best Accuracy: %.2f%% (Epoch %d)\n", bestAccuracy*100, bestEpoch)

	// Test after training with debug info
	fmt.Println("\nShape recognition predictions after training:")
	shapeNames := []string{"Circle", "Square", "Triangle", "Rectangle", "Oval", "Star", "Heart", "Pentagon"}
	confusionMatrix := make([][]int, numClasses)
	for i := range confusionMatrix {
		confusionMatrix[i] = make([]int, numClasses)
	}

	// Test on unique patterns
	testInputs := [][]bool{
		{false, false, true}, // Circle
		{true, true, true},   // Square
		{true, true, false},  // Triangle
		{true, true, true},   // Rectangle
		{false, false, true}, // Oval
		{true, true, false},  // Star
		{false, false, true}, // Heart
		{true, true, true},   // Pentagon
	}
	testTargets := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for i, input := range testInputs {
		prediction := model.Predict(input)
		confusionMatrix[testTargets[i]][prediction]++
		fmt.Printf("Input %v (Features: %s): predicted=%s, target=%s\n",
			input,
			describeFeatures(input),
			shapeNames[prediction],
			shapeNames[testTargets[i]])
	}

	// Print confusion matrix
	fmt.Println("\nConfusion Matrix:")
	fmt.Println("Predicted →")
	fmt.Println("Actual ↓")
	fmt.Print("    ")
	for i := 0; i < numClasses; i++ {
		fmt.Printf("%-8s", shapeNames[i][:4])
	}
	fmt.Println()
	for i := 0; i < numClasses; i++ {
		fmt.Printf("%-4s", shapeNames[i][:4])
		for j := 0; j < numClasses; j++ {
			fmt.Printf("%-8d", confusionMatrix[i][j])
		}
		fmt.Println()
	}

	// Calculate and print detailed metrics
	fmt.Println("\nDetailed Metrics:")
	fmt.Printf("%-10s %-10s %-10s %-10s %-10s\n", "Class", "Precision", "Recall", "F1-Score", "Support")
	fmt.Println(strings.Repeat("-", 50))

	for i := 0; i < numClasses; i++ {
		precision := calculatePrecision(confusionMatrix, i)
		recall := calculateRecall(confusionMatrix, i)
		f1Score := calculateF1Score(precision, recall)
		support := calculateSupport(confusionMatrix, i)

		fmt.Printf("%-10s %-10.2f %-10.2f %-10.2f %-10d\n",
			shapeNames[i],
			precision*100,
			recall*100,
			f1Score*100,
			support)
	}

	// Calculate macro and micro averages
	macroPrecision, macroRecall, macroF1 := calculateMacroAverages(confusionMatrix)
	microPrecision, microRecall, microF1 := calculateMicroAverages(confusionMatrix)

	fmt.Println("\nOverall Metrics:")
	fmt.Printf("Macro Average - Precision: %.2f%%, Recall: %.2f%%, F1-Score: %.2f%%\n",
		macroPrecision*100, macroRecall*100, macroF1*100)
	fmt.Printf("Micro Average - Precision: %.2f%%, Recall: %.2f%%, F1-Score: %.2f%%\n",
		microPrecision*100, microRecall*100, microF1*100)

	// Test with variations
	fmt.Println("\nTesting with variations:")
	variationInputs := [][]bool{
		{false, false, false}, // Almost circle but asymmetric
		{true, true, false},   // Almost square but asymmetric
		{false, true, true},   // Has corners but curved
	}
	for _, input := range variationInputs {
		prediction := model.Predict(input)
		fmt.Printf("Input %v (Features: %s): predicted=%s\n",
			input,
			describeFeatures(input),
			shapeNames[prediction])
	}
}

func calculateAccuracy(model *tmffi.Model, inputs [][]bool, targets []int) float64 {
	correct := 0
	for i, input := range inputs {
		prediction := model.Predict(input)
		if prediction == targets[i] {
			correct++
		}
	}
	return float64(correct) / float64(len(inputs))
}

func calculatePrecision(confusionMatrix [][]int, class int) float64 {
	totalPredicted := 0
	for i := 0; i < len(confusionMatrix); i++ {
		totalPredicted += confusionMatrix[i][class]
	}
	if totalPredicted == 0 {
		return 0
	}
	return float64(confusionMatrix[class][class]) / float64(totalPredicted)
}

func calculateRecall(confusionMatrix [][]int, class int) float64 {
	totalActual := 0
	for j := 0; j < len(confusionMatrix); j++ {
		totalActual += confusionMatrix[class][j]
	}
	if totalActual == 0 {
		return 0
	}
	return float64(confusionMatrix[class][class]) / float64(totalActual)
}

func calculateF1Score(precision, recall float64) float64 {
	if precision+recall == 0 {
		return 0
	}
	return 2 * (precision * recall) / (precision + recall)
}

func calculateSupport(confusionMatrix [][]int, class int) int {
	total := 0
	for j := 0; j < len(confusionMatrix); j++ {
		total += confusionMatrix[class][j]
	}
	return total
}

func calculateMacroAverages(confusionMatrix [][]int) (float64, float64, float64) {
	var totalPrecision, totalRecall, totalF1 float64
	numClasses := len(confusionMatrix)

	for i := 0; i < numClasses; i++ {
		precision := calculatePrecision(confusionMatrix, i)
		recall := calculateRecall(confusionMatrix, i)
		f1 := calculateF1Score(precision, recall)

		totalPrecision += precision
		totalRecall += recall
		totalF1 += f1
	}

	return totalPrecision / float64(numClasses),
		totalRecall / float64(numClasses),
		totalF1 / float64(numClasses)
}

func calculateMicroAverages(confusionMatrix [][]int) (float64, float64, float64) {
	var totalTP, totalFP, totalFN float64
	numClasses := len(confusionMatrix)

	for i := 0; i < numClasses; i++ {
		totalTP += float64(confusionMatrix[i][i])
		for j := 0; j < numClasses; j++ {
			if i != j {
				totalFP += float64(confusionMatrix[j][i])
				totalFN += float64(confusionMatrix[i][j])
			}
		}
	}

	precision := totalTP / (totalTP + totalFP)
	recall := totalTP / (totalTP + totalFN)
	f1 := calculateF1Score(precision, recall)

	return precision, recall, f1
}

func describeFeatures(input []bool) string {
	features := []string{}
	if input[0] {
		features = append(features, "straight")
	} else {
		features = append(features, "curved")
	}
	if input[1] {
		features = append(features, "corners")
	} else {
		features = append(features, "no corners")
	}
	if input[2] {
		features = append(features, "symmetric")
	} else {
		features = append(features, "asymmetric")
	}
	return fmt.Sprintf("%v", features)
}
