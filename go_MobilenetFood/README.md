# WasmEdge-Go Tensorflow Extension MobileNet-Food example

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

The pre-built WASM from rust is provided as "rust_mobilenet_food_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_mobilenet_food
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_mobilenet_food_lib.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./mobilenet_food rust_mobilenet_food_lib.wasm food.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mobilenet_food rust_mobilenet_food_lib.wasm.so food.jpg]
RUST: Loaded image in ... 8.464961ms
RUST: Resized image in ... 10.845873ms
RUST: Parsed output in ... 237.867977ms
RUST: index 258, prob 0.7628046
RUST: Finished post-processing in ... 237.983011ms
GO: Run bindgen -- infer: It is very likely a <a href='https://www.google.com/search?q=Hot dog'>Hot dog</a> in the picture
```
