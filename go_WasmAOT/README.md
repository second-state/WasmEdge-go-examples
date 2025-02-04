# WasmEdge-Go WASM AOT example

This example provide a tool for compiling a WASM file into compiled-WASM for AOT mode.

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/docs/start/install) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.13.5
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.13.5
$ go build
```

## Run

```bash
./wasmAOT fibonacci.wasm fibonacci.wasm.so
```

The output will be as the following:

```bash
Go: Args: [./wasmAOT fibonacci.wasm fibonacci.wasm.so]
```
