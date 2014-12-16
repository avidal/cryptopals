all: clean go rust
.PHONY: all go rust

go:
	cd go && go test -v ./...

rust:
	cd rust && cargo test --verbose

clean:
	cd rust && cargo clean
	cd go && go clean

