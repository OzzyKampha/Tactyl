package tmffi

/*
#cgo LDFLAGS: -L${SRCDIR}/../../Tactyl-engine/target/release -ltactyl_engine
#include <stdint.h>

int32_t predict(const char* bit_vector, size_t len);
*/
import "C"
import "unsafe"

// Predict calls the Rust predict function with a slice of binary values (0 or 1).
func Predict(bitVector []uint8) int32 {
	if len(bitVector) == 0 {
		return 0
	}
	return int32(C.predict((*C.char)(unsafe.Pointer(&bitVector[0])), C.size_t(len(bitVector))))
}
