# WasmEdge-Go example for passing string using memory pointer

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
tinygo build -o greet.wasm -target wasi greet.go

go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
go build greet_memory.go
```

## Run

```bash
# Run in interpreter mode
./greet_memory greet.wasm
```

The output will be as the following:

```bash
Hello, WasmEdge!
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.
