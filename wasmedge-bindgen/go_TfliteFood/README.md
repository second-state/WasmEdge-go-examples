# WasmEdge-Go Tensorflow Extension MobileNet example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) with the `TensorFlow` extension is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -e tf -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
$ go get github.com/second-state/wasmedge-bindgen@v0.4.1
$ go build -tags tensorflow
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_tflite_food_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_tflite_food
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_tflite_food_lib.wasm`.
```

## Run

```bash
# Run in interpreter mode
$ ./tflite_food rust_tflite_food_lib.wasm food.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./tflite_food rust_tflite_food_lib.wasm food.jpg]
It is very likely a Hot dog in the picture
```

If you want to try this example in AOT mode, you can run the following command or follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

```bash
wasmedgec rust_tflite_food_lib.wasm rust_tflite_food_lib.so
./tflite_food rust_tflite_food_lib.so food.jpg
```

And you will notice how big the performance will be enhanced.
