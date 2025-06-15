package tmffi

/*
#cgo LDFLAGS: -L${SRCDIR}/../../Tactyl-engine/target/release -ltactyl_engine
#include <stdint.h>
#include <stdlib.h>

struct TactylModel;
typedef struct TactylModel TactylModel;

TactylModel* tm_create_model(size_t num_features, size_t num_clauses, size_t num_classes, int32_t t, double s);
void tm_free_model(TactylModel* ptr);
void tm_train(TactylModel* ptr, uint64_t input, int32_t target);
int32_t tm_predict(const TactylModel* ptr, uint64_t input);
int32_t tm_predict_debug(const TactylModel* ptr, uint64_t input);
void tm_print_state_histogram(const TactylModel* ptr);
*/
import "C"

type Model struct {
	ptr *C.TactylModel
}

func NewModel(numFeatures, numClauses, numClasses, t int, s float64) *Model {
	return &Model{ptr: C.tm_create_model(C.size_t(numFeatures), C.size_t(numClauses), C.size_t(numClasses), C.int32_t(t), C.double(s))}
}

func (m *Model) Train(input uint64, target int32) {
	C.tm_train(m.ptr, C.uint64_t(input), C.int32_t(target))
}

func (m *Model) Predict(input uint64) int {
	return int(C.tm_predict(m.ptr, C.uint64_t(input)))
}

func (m *Model) PredictDebug(input uint64) int {
	return int(C.tm_predict_debug(m.ptr, C.uint64_t(input)))
}

func (m *Model) Free() {
	C.tm_free_model(m.ptr)
	m.ptr = nil
}

func (m *Model) PrintStateHistogram() {
	C.tm_print_state_histogram(m.ptr)
}
