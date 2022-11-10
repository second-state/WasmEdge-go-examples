# WasmEdge-Go Wasm-Bindgen Functions example

This example is a rust to WASM with `wasmedge-bindgen`. This example is modified from the [nodejs WASM example](https://github.com/second-state/wasm-learning/tree/master/nodejs/functions).

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
$ go get github.com/second-state/wasmedge-bindgen@v0.4.1
$ go build
```

## (Optional) Build the example WASM from rust

The pre-built WASM from rust is provided as "rust_bindgen_funcs_lib.wasm".

For building the WASM from the rust source, the following steps are required:

* Install the [rustc and cargo](https://www.rust-lang.org/tools/install).
* Install the target `wasm32-wasi` for rust.
  * `$ rustup target add wasm32-wasi`

```bash
$ cd rust_bindgen_funcs
$ cargo build --target wasm32-wasi --release
# The output WASM will be `target/wasm32-wasi/release/rust_bindgen_funcs_lib.wasm`.
```

## Run

```bash
# Run in interpreter mode
$ ./bindgen_funcs rust_bindgen_funcs_lib.wasm
```

The standard output of this example will be the following:

```bash
Run bindgen -- create_line: {"points":[{"x":2.5,"y":7.8},{"x":2.5,"y":5.8}],"valid":true,"length":2.0,"desc":"A thin red line"}
Run bindgen -- say: hello bindgen funcs test
Run bindgen -- obfusticate: N dhvpx oebja sbk whzcf bire gur ynml qbt
Run bindgen -- lowest_common_multiple: 246
Run bindgen -- sha3_digest: [87 27 231 209 189 105 251 49 159 10 211 250 15 159 154 181 43 218 26 141 56 199 25 45 60 10 20 163 54 211 195 203]
Run bindgen -- keccak_digest: [126 194 241 200 151 116 227 33 216 99 159 22 107 3 177 169 216 191 114 156 174 193 32 159 246 228 245 133 52 75 55 27]
```

If you want to try this example in AOT mode, you can run the following command or follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

```bash
wasmedgec rust_bindgen_funcs_lib.wasm rust_bindgen_funcs_lib.so
./bindgen_funcs rust_bindgen_funcs_lib.so
```

And you will notice how big the performance will be enhanced.
