# WasmEdge-Go Read file example

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

The pre-built WASM from rust is provided as "rust_readfile.wasm".

For building the WASM from the rust source, the following steps are required:

```bash
$ cd rust_readfile
$ rustwasmc build
# The output file will be `pkg/rust_readfile.wasm`.
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.

## Run

```bash
# Run in interpreter mode
$ ./read_file rust_readfile.wasm file.txt
```

The standard output of this example will be the following:

```bash
Rust: Opening input file "file.txt"...
Rust: Read input file "file.txt" succeeded.
Rust: Please input the line number to print the line of file.
```

Then the developers can interact with this program with standard input:

```bash
# Input "5" and press Enter.
5
# The output will be the 5th line of `file.txt`:
abcDEF___!@#$%^
# Input "-3" and press Enter.
-3
# The output will specify the error:
Rust: ERROR - Input "-3" is not an integer: invalid digit found in string
# Input "15" and press Enter.
15
# The output will specify that the file is smaller than 15 lines:
Rust: ERROR - Line "15" is out of range.
# Input "abcd" and press Enter.
abcd
# The output will specify the error:
Rust: ERROR - Input "abcd" is not an integer: invalid digit found in string
# To terminate the program, send the EOF (Ctrl + D).
^D
# The output will print the terminate message:
Rust: Process end.
```
