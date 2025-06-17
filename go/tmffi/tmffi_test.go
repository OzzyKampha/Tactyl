package tmffi

import (
	"testing"
)

func TestTsetlinMachine(t *testing.T) {
	// Create a Tsetlin Machine with:
	// - 2 input features
	// - 2 output classes
	// - s = 3.9 (specificity)
	// - 10 clauses per class
	// - 100 automata states
	// - No dropout
	// - T = 10 (learning threshold)
	tm := NewModel(2, 2, 3.9, 10, 10, 100, 0.0)
	if tm == nil {
		t.Fatal("Failed to create Tsetlin Machine")
	}
	defer tm.Free()

	// Train on a sample
	input := []bool{true, false}
	tm.Train(input, 0) // train for class 0

	// Predict
	prediction := tm.Predict(input)
	if prediction != 0 {
		t.Errorf("Expected prediction 0, got %d", prediction)
	}
}
