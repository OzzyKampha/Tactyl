package main

import (
	"fmt"
	"tactyl/tmffi"
)

func main() {
	numFeatures := 4
	model := tmffi.NewModel(numFeatures)
	defer model.Free()

	inputs := []uint64{
		0b0011,
		0b0101,
		0b1111,
	}
	feedbacks := []int32{1, 2, 1}

	for i, input := range inputs {
		fmt.Printf("\n--- Training step %d: input = %04b, feedback = %d ---\n", i+1, input, feedbacks[i])
		model.Train(input, feedbacks[i])
	}

	score := model.Predict(inputs[len(inputs)-1])
	fmt.Println("Prediction score:", score)
}
