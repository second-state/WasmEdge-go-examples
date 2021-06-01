# WasmEdge-Go Wasm-Bindgen Kmeans example

This example is a rust to WASM with `wasm-bindgen`. This example is modified from the [nodejs WASM example](https://github.com/second-state/wasm-learning/tree/master/nodejs/kmeans).

## Build

```bash
# In the current directory.
$ go get -u github.com/second-state/WasmEdge-go
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_bindgen_kmeans_lib_bg.wasm".
The pre-built compiled-WASM from rust is provided as "rust_bindgen_kmeans_lib_bg.wasm.so".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Set the default `rustup` version to `1.50.0` or lower.
  * `$ rustup default 1.50.0`
* Install the [rustwasmc](https://github.com/second-state/rustwasmc)
  * `$ curl https://raw.githubusercontent.com/second-state/rustwasmc/master/installer/init.sh -sSf | sh`

```bash
$ cd rust_bindgen_kmeans
$ rustwasmc build --enable-aot
# The output WASM will be `pkg/rust_bindgen_kmeans_lib_bg.wasm`.
```

For compiling the WASM to SO for the AOT mode, please follow the tools of [WasmEdge](https://github.com/WasmEdge/WasmEdge):

```bash
$ wget https://github.com/WasmEdge/WasmEdge/releases/download/0.8.0/WasmEdge-0.8.0-manylinux2014_x86_64.tar.gz
$ tar -xzf WasmEdge-0.8.0-manylinux2014_x86_64.tar.gz
$ ./WasmEdge-0.8.0-Linux/bin/wasmedgec rust_bindgen_kmeans_lib_bg.wasm rust_bindgen_kmeans_lib_bg.wasm.so
# The output compiled-WASM will be at `rust_bindgen_kmeans_lib_bg.wasm.so`.
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
# For the interpreter mode
$ ./bindgen_kmeans rust_bindgen_kmeans_lib_bg.wasm
# For the AOT mode
$ ./bindgen_kmeans rust_bindgen_kmeans_lib_bg.wasm.so
```

The standard output of this example will be the following:

```bash
```