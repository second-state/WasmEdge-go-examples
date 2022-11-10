# WasmEdge-Go Accumulate static variable example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "accumulate_bg.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc](https://www.rust-lang.org/tools/install).
* Install the `wasm32-wasi` target: `$ rustup target add wasm32-wasi`
* Install the [rustwasmc](https://github.com/second-state/rustwasmc).

```bash
$ rustwasmc build
# The output file will be `pkg/accumulate_bg.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./accumulate accumulate_bg.wasm
```

The standard output of this example will be the following:

```bash
1
2
3
```
