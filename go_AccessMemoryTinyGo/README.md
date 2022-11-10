# WasmEdge-Go memory access example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
tinygo build -o fib.wasm -target wasi fib.go
# Test using wasmedge command
wasmedge fib.wasm

go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
go build run.go
```

## Run

```bash
# Run in interpreter mode
./run fib.wasm
```

The output will be as the following:

```bash
fibArray() returned: 34
fibArray memory at: 0x14d3c
fibArray: [0 0 0 0 1 0 0 0 1 0 0 0 2 0 0 0 3 0 0 0 5 0 0 0 8 0 0 0 13 0 0 0 21 0 0 0 34 0 0 0]
fibArrayReturnMemory memory at: 0x14d4c
fibArrayReturnMemory: [0 0 0 0 1 0 0 0 1 0 0 0 2 0 0 0 3 0 0 0 5 0 0 0 8 0 0 0 13 0 0 0 21 0 0 0 34 0 0 0]
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.
