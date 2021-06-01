# WasmEdge-Go WASM AOT example

This example provide a tool for compiling a WASM file into compiled-WASM for AOT mode.

## Build

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
2021-06-01 13:59:18,405 INFO [default] compile start
2021-06-01 13:59:18,406 INFO [default] verify start
2021-06-01 13:59:18,406 INFO [default] optimize start
2021-06-01 13:59:18,411 INFO [default] codegen start
2021-06-01 13:59:18,435 INFO [default] compile done
```
