# WasmEdge-Go external reference example

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

## WASM file introduction

The WASM file is converted from WAT:

```wasm
(module
  (type $t0 (func (param externref i32) (result i32)))
  (type $t1 (func (param externref i32 i32) (result i32)))
  (type $t2 (func (param externref externref i32 i32) (result i32)))
  (import "extern_module" "square" (func $square (type $t0)))
  (import "extern_module" "add" (func $add (type $t1)))
  (import "extern_module" "mul" (func $mul (type $t1)))
  (func $call_add (export "call_add") (type $t1) (param $p0 externref) (param $p1 i32) (param $p2 i32) (result i32)
    (call $add
      (local.get $p0)
      (local.get $p1)
      (local.get $p2)))
  (func $call_mul (export "call_mul") (type $t1) (param $p0 externref) (param $p1 i32) (param $p2 i32) (result i32)
    (call $mul
      (local.get $p0)
      (local.get $p1)
      (local.get $p2)))
  (func $call_square (export "call_square") (type $t0) (param $p0 externref) (param $p1 i32) (result i32)
    (call $square
      (local.get $p0)
      (local.get $p1)))
  (func $call_add_square (export "call_add_square") (type $t2) (param $p0 externref) (param $p1 externref) (param $p2 i32) (param $p3 i32) (result i32)
    (call $square
      (local.get $p1)
      (call $add
        (local.get $p0)
        (local.get $p2)
        (local.get $p3))))
  (memory $memory (export "memory") 1))
```

This WASM exported 4 functions:

* `call_add`
  * Params {`externref`, `i32`, `i32`}, Returns {`i32`}
  * Call the host function `extern_module::add`, and the real add function implementation is passed in by the `externref`.
* `call_mul`:
  * Params {`externref`, `i32`, `i32`}, Returns {`i32`}
  * Call the host function `extern_module::mul`, and the real mul function implementation is passed in by the `externref`.
* `call_square`:
  * Params {`externref`, `i32`}, Returns {`i32`}
  * Call the host function `extern_module::square`, and the real square function implementation is passed in by the `externref`.
* `call_add_square`:
  * Params {`externref`, `externref`, `i32`, `i32`}, Returns {`i32`}
  * Call the host function `extern_module::add`, and the real add function implementation is passed in by the first `externref` argument.
  * And then call the host function `extern_module::square`, and the real square function implementation is passed in by the second `externref` argument.

## Run

```bash
./externref funcs.wasm
```

The standard output of this example will be the following:

```bash
Go: Args: [./externref funcs.wasm]
Go: Entering go host function host_add
Go: Entering go function real_add
Go: Leaving go function real_add
Go: Leaving go host function host_add
Run call_add: 1234 + 5678 = 6912
Go: Entering go host function host_mul
Go: Entering go function real_mul
Go: Leaving go function real_mul
Go: Leaving go host function host_mul
Run call_mul: 4827 * (-31519) = -152142213
Go: Entering go host function host_square
Go: Entering go function real_square
Go: Leaving go function real_square
Go: Leaving go host function host_square
Run call_square: 1024^2 = 1048576
Go: Entering go host function host_add
Go: Entering go function real_add
Go: Leaving go function real_add
Go: Leaving go host function host_add
Go: Entering go host function host_square
Go: Entering go function real_square
Go: Leaving go function real_square
Go: Leaving go host function host_square
Run call_square: (761 + 195)^2 = 913936
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.
