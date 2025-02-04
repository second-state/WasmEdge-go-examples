# WasmEdge-Go Tensorflow Extension MTCNN example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/docs/start/install) with the `WasmEdge-TensorFlow` plug-in is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- --plugins wasmedge_tensorflow  -v 0.13.5
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.13.5
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_mtcnn.wasm".
The pre-built and compiled WASM for AOT mode is provided as "rust_mtcnn.wasm.so".

For building the WASM from the rust source, the following steps are required:

* Install the [rust and cargo](https://www.rust-lang.org/tools/install).
* Install the `wasm32-wasi` target: `$ rustup target add wasm32-wasi`

```bash
$ cd rust_mtcnn
$ cargo build --release --target=wasm32-wasi
# The output wasm will be at `target/wasm32-wasi/release/rust_mtcnn.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./mtcnn rust_mtcnn.wasm mtcnn.pb solvay.jpg out.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mtcnn rust_mtcnn.wasm mtcnn.pb solvay.jpg out.jpg]
Drawing box: 30 results ...
```

And the output image will be at `out.jpg`.
