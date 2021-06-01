# WasmEdge-Go Wasm-Bindgen Functions example

This example is a rust to WASM with `wasm-bindgen`. This example is modified from the [nodejs WASM example](https://github.com/second-state/wasm-learning/tree/master/nodejs/functions).

## Build

```bash
# In the current directory.
$ go get -u github.com/second-state/WasmEdge-go
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_bindgen_funcs_lib_bg.wasm".
The pre-built compiled-WASM from rust is provided as "rust_bindgen_funcs_lib_bg.so".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Set the default `rustup` version to `1.50.0` or lower.
  * `$ rustup default 1.50.0`
* Install the [rustwasmc](https://github.com/second-state/rustwasmc)
  * `$ curl https://raw.githubusercontent.com/second-state/rustwasmc/master/installer/init.sh -sSf | sh`

```bash
$ cd rust_bindgen_funcs
$ rustwasmc build
# The output WASM will be `pkg/rust_bindgen_funcs_lib_bg.wasm`.
```

For compiling the WASM to SO for the AOT mode, please follow the tools of [WasmEdge](https://github.com/WasmEdge/WasmEdge):

```bash
$ wget https://github.com/WasmEdge/WasmEdge/releases/download/0.8.0/WasmEdge-0.8.0-manylinux2014_x86_64.tar.gz
$ tar -xzf WasmEdge-0.8.0-manylinux2014_x86_64.tar.gz
$ ./WasmEdge-0.8.0-Linux/bin/wasmedgec rust_bindgen_funcs_lib_bg.wasm rust_bindgen_funcs_lib_bg.wasm.so
# The output compiled-WASM will be at `rust_bindgen_funcs_lib_bg.wasm.so`.
```

Or follow the [example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) for compiling the WASM to SO:

```bash
# In the `go_WasmAOT` directory
$ go get -u .
$ go build
# Prepare the input WASM file
$ ./wasmAOT input.wasm output.wasm.so
```

## Run

```bash
# For the interpreter mode
$ ./bindgen_funcs rust_bindgen_funcs_lib_bg.wasm
# For the AOT mode
$ ./bindgen_funcs rust_bindgen_funcs_lib_bg.so
```

The standard output of this example will be the following:

```bash
Run bindgen -- create_line: {"points":[{"x":1.5,"y":3.8},{"x":2.5,"y":5.8}],"valid":true,"length":2.2360682,"desc":"A thin red line"}
Run bindgen -- say: hello bindgen funcs test
Run bindgen -- obfusticate: N dhvpx oebja sbk whzcf bire gur ynml qbt
Run bindgen -- lowest_common_multiple: 246
Run bindgen -- sha3_digest: [87 27 231 209 189 105 251 49 159 10 211 250 15 159 154 181 43 218 26 141 56 199 25 45 60 10 20 163 54 211 195 203]
Run bindgen -- keccak_digest: [126 194 241 200 151 116 227 33 216 99 159 22 107 3 177 169 216 191 114 156 174 193 32 159 246 228 245 133 52 75 55 27]
```