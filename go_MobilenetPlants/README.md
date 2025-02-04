# WasmEdge-Go Tensorflow Extension MobileNet-Plants example

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

The pre-built WASM from rust is provided as "rust_mobilenet_plants_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_mobilenet_plants
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_mobilenet_plants_lib.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./mobilenet_plants rust_mobilenet_plants_lib.wasm sunflower.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mobilenet_plants rust_mobilenet_plants_lib.wasm.so sunflower.jpg]
RUST: Loaded image in ... 12.267229ms
RUST: Resized image in ... 14.206377ms
RUST: Parsed output in ... 294.487086ms
RUST: index 1680, prob 0.9295
RUST: Finished post-processing in ... 294.701533ms
GO: Run bindgen -- infer: It is very likely a <a href='https://www.google.com/search?q=Helianthus annuus'>Helianthus annuus</a> in the picture
```
