# WasmEdge-Go Tensorflow Extension MobileNet-Insects example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/docs/start/install) with the `WasmEdge-TensorFlow` plug-in is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- --plugins wasmedge_tensorflow  -v 0.13.0
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.13.0
$ go get github.com/second-state/wasmedge-bindgen@v0.4.1
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_mobilenet_insects_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_mobilenet_insects
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_mobilenet_insects_lib.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./mobilenet_insects rust_mobilenet_insects_lib.wasm ladybug.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mobilenet_insects rust_mobilenet_insects_lib.wasm.so ladybug.jpg]
RUST: Loaded image in ... 5.628365ms
RUST: Resized image in ... 7.165381ms
RUST: Parsed output in ... 270.750735ms
RUST: index 320, prob 0.92208606
RUST: Finished post-processing in ... 271.020909ms
GO: Run bindgen -- infer: It is very likely a <a href='https://www.google.com/search?q=Coccinella septempunctata'>Coccinella septempunctata</a> in the picture
```
