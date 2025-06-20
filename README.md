# Tactyl

A fast, modular, and embeddable [Tsetlin Machine](https://en.wikipedia.org/wiki/Tsetlin_machine) engine written in Rust, with Go bindings.

## Features

- **Core TM Features:**
  - Binary Tsetlin Machine implementation
  - Two automata per feature per clause (include/exclude)
  - Clause polarity (positive/negative)
  - Stochastic feedback with configurable threshold (T) and specificity (s)
  - Debugging tools for automata state and clause firing

- **Implementation Variants:**
  - Standard Tsetlin Machine implementation
  - Parallel processing capabilities
  - Memory-efficient storage options

- **Integration:**
  - Exposes a C-compatible API for integration with other languages (e.g., Go, Python, C)
  - Rust-native API with full type safety

## Getting Started

### Usage

You can use the engine as a Rust library or via FFI from other languages.

#### Example: Creating and Training a Model (Rust)

```rust
use tactyl_engine::*;

// Standard TM
let mut standard_tm = SharedTsetlinMachine::new(4, 2, 3.9, 10, 100, 0.1, 15);

// Train on data
let input = vec![true, false, true, false];
standard_tm.train(&input, 1);

// Predict
let prediction = standard_tm.predict(&input);
```

#### C FFI

All main functions are exposed as `extern "C"` for use in C, Go, Python, etc.

- `tm_create_model(num_features, num_clauses, t, s) -> *mut TactylModel`
- `tm_train(model_ptr, input, target)`
- `tm_predict(model_ptr, input) -> i32`
- `tm_free_model(model_ptr)`
- `tm_print_state_histogram(model_ptr)`

### Performance Example

Here's an example showing the training and prediction speeds for different TM implementations:

```rust
use tactyl_engine::*;
use std::time::Instant;

// Create different TM implementations
let mut standard_tm = SharedTsetlinMachine::new(8, 2, 3.9, 50, 100, 0.1, 15);
let mut parallel_tm = ParallelBitwiseTsetlinMachine::new(8, 2, 3.9, 50, 100, 0.1, 15);
let mut bitpacked_tm = BitpackedTsetlinMachine::new(8, 2, 3.9, 50, 100, 0.1, 15);

// Generate synthetic training data
let training_data: Vec<(Vec<bool>, usize)> = (0..100)
    .map(|i| {
        let input = (0..8).map(|b| (i & (1 << b)) != 0).collect::<Vec<_>>();
        let target = if i % 2 == 0 { 0 } else { 1 };
        (input, target)
    })
    .collect();

// Measure training time
let start = Instant::now();
for (input, target) in &training_data {
    standard_tm.train(input, *target);
}
let standard_train_time = start.elapsed();

let start = Instant::now();
for (input, target) in &training_data {
    parallel_tm.train(input, *target);
}
let parallel_train_time = start.elapsed();

let start = Instant::now();
for (input, target) in &training_data {
    bitpacked_tm.train(input, *target);
}
let bitpacked_train_time = start.elapsed();

println!("Training Times:");
println!("  Standard TM:   {:?}", standard_train_time);
println!("  Parallel TM:   {:?}", parallel_train_time);
println!("  Bitpacked TM:  {:?}", bitpacked_train_time);

// Measure prediction time
let test_data: Vec<Vec<bool>> = (0..50)
    .map(|i| (0..8).map(|b| (i & (1 << b)) != 0).collect::<Vec<_>>())
    .collect();

let start = Instant::now();
for input in &test_data {
    standard_tm.predict(input);
}
let standard_pred_time = start.elapsed();

let start = Instant::now();
for input in &test_data {
    parallel_tm.predict(input);
}
let parallel_pred_time = start.elapsed();

let start = Instant::now();
for input in &test_data {
    bitpacked_tm.predict(input);
}
let bitpacked_pred_time = start.elapsed();

println!("Prediction Times:");
println!("  Standard TM:   {:?}", standard_pred_time);
println!("  Parallel TM:   {:?}", parallel_pred_time);
println!("  Bitpacked TM:  {:?}", bitpacked_pred_time);
```

### Debugging

- Use `tm_print_state_histogram(model_ptr)` to print automata state distributions
- The engine prints detailed debug info during training if compiled with debug output enabled

## License

Tactyl is open core. See individual components for their respective licenses.

## References

- [Tsetlin Machine Paper](https://arxiv.org/abs/1804.01508)
- [PyTsetlinMachine](https://github.com/cair/TsetlinMachine)

---

*This project is under active development.* 