# WasmEdge-Go example for passing string using memory pointer

## Build

Before trying this example, the [WasmEdge installation](https://github.com/WasmEdge/WasmEdge/blob/master/docs/install.md) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.9.0-rc.5
```

Then you can build this example.

```bash
go get github.com/second-state/WasmEdge-go/wasmedge@v0.9.0-rc5
go build greet_memory.go
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_memory_greet_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install `wasm32-wasi` target
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_memory_greet
$ cargo build --target wasm32-wasi
# The output WASM will be `target/wasm32-wasi/debug/rust_memory_greet_lib.wasm`.
```

## Run

```bash
# Run in interpreter mode
./memory-greet rust_memory_greet_lib.wasm
```

The output will be as the following:
```
Hello, WasmEdge!
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.
