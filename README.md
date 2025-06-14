# Tactyl

A fast, modular, and embeddable [Tsetlin Machine](https://en.wikipedia.org/wiki/Tsetlin_machine) engine written in Rust, with Go bindings.

## Features

- Binary Tsetlin Machine implementation
- Two automata per feature per clause (include/exclude)
- Clause polarity (positive/negative)
- Stochastic feedback with configurable threshold (T) and specificity (s)
- Debugging tools for automata state and clause firing
- Exposes a C-compatible API for integration with other languages (e.g., Go, Python, C)

## Getting Started

### Usage

You can use the engine as a Rust library or via FFI from other languages.

#### Example: Creating and Training a Model (Rust)

```rust
use tactyl_engine::*;

let num_features = 2;
let num_clauses = 10;
let t = 5;
let s = 3.9;

let model = tm_create_model(num_features, num_clauses, t, s);
// Train on a sample (input as u64, target as i32)
tm_train(model, 0b11, 1);
// Predict
let prediction = tm_predict(model, 0b11);
tm_free_model(model);
```

#### C FFI

All main functions are exposed as `extern "C"` for use in C, Go, Python, etc.

- `tm_create_model(num_features, num_clauses, t, s) -> *mut TactylModel`
- `tm_train(model_ptr, input, target)`
- `tm_predict(model_ptr, input) -> i32`
- `tm_free_model(model_ptr)`
- `tm_print_state_histogram(model_ptr)`

### Parameters

- `num_features`: Number of input features (max 64 for now, packed in a `u64`)
- `num_clauses`: Number of clauses (should be even, half positive, half negative)
- `t`: Threshold for voting (higher = more consensus needed)
- `s`: Specificity (controls feedback probability, higher = more specific clauses)

### Debugging

- Use `tm_print_state_histogram(model_ptr)` to print automata state distributions.
- The engine prints detailed debug info during training if compiled with debug output enabled.

## License

Tactyl is open core. Tactyl-engine is protected IP and not open source. All rights reserved.

## References

- [Tsetlin Machine Paper](https://arxiv.org/abs/1804.01508)
- [PyTsetlinMachine](https://github.com/cair/TsetlinMachine)

---

*This project is under active development.* 