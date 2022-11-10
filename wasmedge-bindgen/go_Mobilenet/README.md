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
["military uniform","medium"]
```

If you want to try this example in AOT mode, you can run the following command or follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

```bash
wasmedgec rust_mobilenet_lib.wasm rust_mobilenet_lib.so
./mobilenet rust_mobilenet_lib.so grace_hopper.jpg
```

And you will notice how big the performance will be enhanced.
