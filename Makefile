.PHONY: all build-rust run-go clean

RUST_LIB_DIR=./Tactyl-engine/target/release
RUST_LIB=$(RUST_LIB_DIR)/libtactyl_engine.so

all: build-rust run-go

build-rust:
	cargo build --release --manifest-path=Tactyl-engine/Cargo.toml
	@test -f $(RUST_LIB) || (echo "Error: $(RUST_LIB) not found" && exit 1)

run-go: build-rust
	cd go && LD_LIBRARY_PATH=$(shell pwd)/Tactyl-engine/target/release go run main.go

clean:
	cargo clean --manifest-path=Tactyl-engine/Cargo.toml 