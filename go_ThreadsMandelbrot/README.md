# WasmEdge-Go Threads Mandelbrot example

This example provide a golang version of [WasmEdge mandelbrot set in threads example](https://github.com/WasmEdge/WasmEdge/tree/master/examples/capi/mandelbrot-set-in-threads).

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

## (Optional) Build the Mondelbrot WASM from C

The pre-built WASM is provided as `mandelbrot.wasm` and can be downloaded here:

```bash
wget https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/examples/capi/mandelbrot-set-in-threads/mandelbrot.wasm
```

For building from C source, please refer to [here](https://github.com/WasmEdge/WasmEdge/tree/master/examples/capi/mandelbrot-set-in-threads#the-mandelbrot-c-program-to-wasm).

## Compile the WASM with AOT Compiler

For better performance, we recommand users to compile the WASM to enable the AOT mode.

```bash
wasmedgec --enable-threads mandelbrot.wasm mandelbrot_aot.wasm
# The `--enable-threads` is needed to turn on the threads proposal supporting.
# The output is in universal WASM format.
# For the shared library format, please use the `.so` on Linux,
# `.dylib` on MacOS, or `.dll` on Windows for the file extension.
```

## (Optional) Output Image Converter Preparation

To convert the output image binary into a PNG file, please refer to the following steps to prepare the converter:

```bash
# Get the converter js code.
wget https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/examples/capi/mandelbrot-set-in-threads/convert.js
# Please check that the `nodejs` and `npm` is installed.
npm install canvas
```

## Run

```bash
# Args: ./threads WASM_PATH [NUM_THREADS]
$ ./threads mandelbrot_aot.wasm 8
Go: Args: [./threads mandelbrot_aot.wasm 8]
Go: Input WASM file: mandelbrot_aot.wasm
Go: Num of threads: 8
Go: Got the result image offset: 1024
```

The output file `output-wasmedge-bin` is the image binary.
Users can use the converter to generate the PNG file:

```bash
node convert.js output-wasmedge-bin output-wasmedge-go.png
```
