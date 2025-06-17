.PHONY: all build-rust run-go clean build-rust-lib copy-lib go-example-multiclass go-examples

RUST_LIB_DIR=./Tactyl-engine/target/release
RUST_LIB=$(RUST_LIB_DIR)/libtactyl_engine.so
GO_EXAMPLE_DIR=go/examples/multiclass

all: build-rust run-go

build-rust:
	cargo build --release --manifest-path=Tactyl-engine/Cargo.toml
	@test -f $(RUST_LIB) || (echo "Error: $(RUST_LIB) not found" && exit 1)

run-go: build-rust
	cd go && LD_LIBRARY_PATH=$(shell pwd)/Tactyl-engine/target/release go run main.go

clean:
	cargo clean --manifest-path=Tactyl-engine/Cargo.toml

build-rust-lib:
	cd Tactyl-engine && cargo build --release --lib

copy-lib: build-rust-lib
	cp $(RUST_LIB) $(GO_EXAMPLE_DIR)/

go-example-multiclass: copy-lib
	cd $(GO_EXAMPLE_DIR) && LD_LIBRARY_PATH=. go run main.go

go-examples: go-example-multiclass 