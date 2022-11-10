# WasmEdge-Go Examples

The [WasmEdge](https://github.com/WasmEdge/WasmEdge) is a high performance WebAssembly runtime optimized for server side applications. The [WasmEdge-Go](https://github.com/second-state/WasmEdge-go) provides a [golang](https://golang.org/) package for accessing to WasmEdge.

## Getting Started

Before trying the examples, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.
The WasmEdge extensions are built for `Linux` platforms. For the `Windows` and `MacOS` users, there is only the WasmEdge without extensions can be installed.

This repository contains examples of WasmEdge-Go.

* Basic examples: The basic API examples for WasmEdge-Go.
  * [go_PrintFibonacci](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_PrintFibonacci): Various WasmEdge-Go API examples in Golang.
  * [go_ReadFile](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_ReadFile): Invoke WASM with WASI in Golang.
  * [go_WasmAOT](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT): WASM AOT compiler in Golang.
  * [go_ExternRef](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_ExternRef): External reference and host functions in Golang examples.
  * [go_ThreadsMandelbrot](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_ThreadsMandelbrot): WASM threads proposal example to draw the mandelbrot image.
* Wasm-Bindgen examples: The tutorials for [wasm-bindgen](https://github.com/rustwasm/wasm-bindgen) from Rust sources and executions.
  * [go_BindgenFuncs](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_BindgenFuncs): Example for invokation the functions from Rust source with wasm-bindgen.
  * [go_BindgenKmeans](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_BindgenKmeans): Example for kmeans calculation with wasm-bindgen.
  * [go_BindgenWasi](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_BindgenWasi): Invoke WASM with WASI and wasm-bindgen.
* WasmEdge-Bindgen examples: The tutorials for [wasmedge-bindgen](https://github.com/second-state/wasmedge-bindgen) from Rust sources and executions.
  * [go_BindgenFuncs](https://github.com/second-state/WasmEdge-go-examples/tree/master/wasmedge-bindgen/go_BindgenFuncs): Example for invokation the functions from Rust source with wasmedge-bindgen.
  * [go_Mobilenet](https://github.com/second-state/WasmEdge-go-examples/tree/master/wasmedge-bindgen/go_Mobilenet): Basic mobilenet example with wasmedge-bindgen.
  * [go_TfliteFood](https://github.com/second-state/WasmEdge-go-examples/tree/master/wasmedge-bindgen/go_TfliteFood): Basic TensorFlow-Lite food example with wasmedge-bindgen..
* Mobilenet examples
  * [go_Mobilenet](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_Mobilenet): Basic mobilenet example with wasm-bindgen.
  * [go_MobilenetBirds](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetBirds), [go_MobilenetFood](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetFood), [go_MobilenetInsects](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetInsects), and [go_MobilenetPlants](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetPlants): MobileNet examples for every categories.
* Tensorflow extension examples
  * [go_mtcnn](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_mtcnn): WasmEdge-tensorflow extension in WasmEdge-Go example.
* Memory examples
  * [go_AccessMemory](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_AccessMemory): Example showing how to pass and return memory pointer in WASM.
  * [go_AccessMemoryTinyGo](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_AccessMemoryTinyGo): Example showing how to pass and return memory pointer in WASM which is compiled from TinyGo.
  * [go_MemoryGreet](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MemoryGreet): Example showing how to pass and return string using pointer in WASM.
  * [go_MemoryGreetTinyGo](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MemoryGreetTinyGo): Example showing how to pass and return string using pointer in WASM which is compiled from TinyGo.
