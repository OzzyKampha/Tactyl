package tmffi

/*
#cgo LDFLAGS: -L${SRCDIR}/../../Tactyl-engine/target/release -ltactyl_engine
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

int32_t predict(const char* bit_vector, size_t len);
void train(const char* bit_vector, size_t len, int32_t target_class);
void set_specificity(double s);
void set_threshold(double threshold);
*/
import "C"
import "unsafe"

// Predict takes a slice of binary values and returns the predicted class
func Predict(input []uint8) int32 {
	if len(input) == 0 {
		return 0
	}

	// Call Rust predict function with direct pointer to input slice
	result := C.predict((*C.char)(unsafe.Pointer(&input[0])), C.size_t(len(input)))
	return int32(result)
}

// Train updates the Tsetlin Machine with a new training example
func Train(input []uint8, targetClass int32) {
	if len(input) == 0 {
		return
	}

	C.train((*C.char)(unsafe.Pointer(&input[0])), C.size_t(len(input)), C.int32_t(targetClass))
}

// SetSpecificity sets the specificity parameter of the Tsetlin Machine
func SetSpecificity(s float64) {
	C.set_specificity(C.double(s))
}

// SetThreshold sets the classification threshold of the Tsetlin Machine
func SetThreshold(threshold float64) {
	C.set_threshold(C.double(threshold))
}
