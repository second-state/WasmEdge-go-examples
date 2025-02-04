# WasmEdge-Go Tensorflow Extension MobileNet example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/docs/start/install) with the `WasmEdge-TensorFlow` plug-in is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- --plugins wasmedge_tensorflow  -v 0.13.5
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.13.5
$ go get github.com/second-state/wasmedge-bindgen@v0.4.1
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_mobilenet_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_mobilenet
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_mobilenet_lib.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./mobilenet rust_mobilenet_lib.wasm grace_hopper.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mobilenet rust_mobilenet_lib.wasm.so grace_hopper.jpg]
RUST: Loaded image in ... 16.522151ms
RUST: Resized image in ... 19.440301ms
RUST: Parsed output in ... 285.83336ms
RUST: index 653, prob 0.43212935
RUST: Finished post-processing in ... 285.995153ms
GO: Run bindgen -- infer: ["military uniform","medium"]
```
