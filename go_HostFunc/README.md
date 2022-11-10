# WasmEdge-Go host function example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_host_func.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_host_func
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_host_func.wasm`.
```

## Run

```bash
./hostfunc rust_host_func.wasm
```

The standard output of this example will be the following:

```bash
Go: Args: [./hostfunc rust_host_func.wasm]
There are 78 'google' in source code of google.com
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.
