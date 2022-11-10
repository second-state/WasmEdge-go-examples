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
$ go build --tags tensorflow
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_mobilenet_lib_bg.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Set the default `rustup` version to `1.50.0` or lower.
  * `$ rustup default 1.50.0`
* Install the [rustwasmc](https://github.com/second-state/rustwasmc)
  * `$ curl https://raw.githubusercontent.com/second-state/rustwasmc/master/installer/init.sh -sSf | sh`

```bash
$ cd rust_mobilenet
$ rustwasmc build
# The output WASM will be `pkg/rust_mobilenet_lib_bg.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./mobilenet rust_mobilenet_lib_bg.wasm grace_hopper.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mobilenet rust_mobilenet_lib_bg.wasm.so grace_hopper.jpg]
RUST: Loaded image in ... 16.522151ms
RUST: Resized image in ... 19.440301ms
RUST: Parsed output in ... 285.83336ms
RUST: index 653, prob 0.43212935
RUST: Finished post-processing in ... 285.995153ms
GO: Run bindgen -- infer: ["military uniform","medium"]
```
