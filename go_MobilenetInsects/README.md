# WasmEdge-Go Tensorflow Extension MobileNet-Insects example

## Build

Before building this project, please ensure the dependency of [WasmEdge-tensorflow extension](https://github.com/second-state/WasmEdge-go#wasmedge-tensorflow-extension) has been installed.

```bash
# In the current directory.
$ go get -u github.com/second-state/WasmEdge-go/wasmedge
$ go build --tags tensorflow
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_mobilenet_insects_lib_bg.wasm".
The pre-built compiled-WASM from rust is provided as "rust_mobilenet_insects_lib_bg.so".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Set the default `rustup` version to `1.50.0` or lower.
  * `$ rustup default 1.50.0`
* Install the [rustwasmc](https://github.com/second-state/rustwasmc)
  * `$ curl https://raw.githubusercontent.com/second-state/rustwasmc/master/installer/init.sh -sSf | sh`

```bash
$ cd rust_mobilenet_insects
$ rustwasmc build
# The output WASM will be `pkg/rust_mobilenet_insects_lib_bg.wasm`.
```

For compiling the WASM to SO for the AOT mode, please follow the tools of [WasmEdge](https://github.com/WasmEdge/WasmEdge):

```bash
$ wget https://github.com/WasmEdge/WasmEdge/releases/download/0.8.0/WasmEdge-0.8.0-manylinux2014_x86_64.tar.gz
$ tar -xzf WasmEdge-0.8.0-manylinux2014_x86_64.tar.gz
$ ./WasmEdge-0.8.0-Linux/bin/wasmedgec rust_mobilenet_insects_lib_bg.wasm rust_mobilenet_insects_lib_bg.wasm.so
# The output compiled-WASM will be at `rust_mobilenet_insects_lib_bg.wasm.so`.
```

Or follow the [example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) for compiling the WASM to SO:

```bash
# In the `go_WasmAOT` directory
$ go get -u github.com/second-state/WasmEdge-go/wasmedge
$ go build
# Prepare the input WASM file
$ ./wasmAOT input.wasm output.wasm.so
```

## Run

```bash
# For interpreter mode:
$ ./mobilenet_insects rust_mobilenet_insects_lib_bg.wasm ladybug.jpg
# For AOT mode:
$ ./mobilenet_insects rust_mobilenet_insects_lib_bg.so ladybug.jpg
```

The standard output of this example will be the following:

```bash
Go: Args: [./mobilenet_insects rust_mobilenet_insects_lib_bg.wasm.so ladybug.jpg]
RUST: Loaded image in ... 5.628365ms
RUST: Resized image in ... 7.165381ms
RUST: Parsed output in ... 270.750735ms
RUST: index 320, prob 0.92208606
RUST: Finished post-processing in ... 271.020909ms
GO: Run bindgen -- infer: It is very likely a <a href='https://www.google.com/search?q=Coccinella septempunctata'>Coccinella septempunctata</a> in the picture
```
