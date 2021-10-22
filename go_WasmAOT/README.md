# WasmEdge-Go WASM AOT example

This example provide a tool for compiling a WASM file into compiled-WASM for AOT mode.

## Build

Before trying this example, the [WasmEdge installation](https://github.com/WasmEdge/WasmEdge/blob/master/docs/install.md) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -p /usr/local
```

Then you can build this example.

```bash
# In the current directory.
$ go get -u github.com/second-state/WasmEdge-go/wasmedge
$ go build
```

## Run

```bash
$ ./wasmAOT fibonacci.wasm fibonacci.wasm.so
```

The output will be as the following:

```
Go: Args: [./wasmAOT fibonacci.wasm fibonacci.wasm.so]
```
