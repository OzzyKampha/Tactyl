package tmffi

/*
#cgo LDFLAGS: -L${SRCDIR}/../../Tactyl-engine/target/release -ltactyl_engine
#include <stdint.h>
#include <stdlib.h>

typedef struct FfiTsetlinMachine FfiTsetlinMachine;

FfiTsetlinMachine* tm_create_model(int input_size, int num_classes, double s, int num_clauses, int t, unsigned char ta_states, float dropout_rate);
void tm_free_model(FfiTsetlinMachine* ptr);
void tm_train_model(FfiTsetlinMachine* ptr, unsigned char* input_ptr, int input_len, int target_class);
int tm_predict_debug(FfiTsetlinMachine* ptr, unsigned char* input_ptr, int input_len);
void tm_print_state_histogram(FfiTsetlinMachine* ptr);
*/
import "C"
import "unsafe"

type Model struct {
	ptr *C.FfiTsetlinMachine
}

func NewModel(inputSize, numClasses int, s float64, numClauses int, t int, taStates uint8, dropoutRate float32) *Model {
	return &Model{ptr: C.tm_create_model(
		C.int(inputSize),
		C.int(numClasses),
		C.double(s),
		C.int(numClauses),
		C.int(t),
		C.uchar(taStates),
		C.float(dropoutRate),
	)}
}

func (m *Model) Train(input []bool, targetClass int) {
	inputLen := len(input)
	inputArr := make([]C.uchar, inputLen)
	for i, b := range input {
		if b {
			inputArr[i] = 1
		} else {
			inputArr[i] = 0
		}
	}
	C.tm_train_model(m.ptr, (*C.uchar)(unsafe.Pointer(&inputArr[0])), C.int(inputLen), C.int(targetClass))
}

func (m *Model) Predict(input []bool) int {
	inputLen := len(input)
	inputArr := make([]C.uchar, inputLen)
	for i, b := range input {
		if b {
			inputArr[i] = 1
		} else {
			inputArr[i] = 0
		}
	}
	return int(C.tm_predict_debug(m.ptr, (*C.uchar)(unsafe.Pointer(&inputArr[0])), C.int(inputLen)))
}

func (m *Model) Free() {
	C.tm_free_model(m.ptr)
	m.ptr = nil
}

func (m *Model) PrintStateHistogram() {
	C.tm_print_state_histogram(m.ptr)
}
