# WasmEdge-Go Examples

The [WasmEdge](https://github.com/WasmEdge/WasmEdge) (formerly SSVM) is a high performance WebAssembly runtime optimized for server side applications. The [WasmEdge-Go](https://github.com/second-state/WasmEdge-go) provides a [golang](https://golang.org/) package for accessing to WasmEdge.

This repository contains examples of WasmEdge-Go.

* Basic examples: The basic API examples for WasmEdge-Go.
  * [go_PrintFibonacci](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_PrintFibonacci): Various WasmEdge-Go API examples in Golang.
  * [go_ReadFile](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_ReadFile): Invoke WASM with WASI in Golang.
  * [go_WasmAOT](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT): WASM AOT compiler in Golang.
  * [go_ExternRef](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_ExternRef): External reference and host functions in Golang examples.
* Wasm-Bindgen examples: The tutorials for [wasm-bindgen](https://github.com/rustwasm/wasm-bindgen) from Rust sources and executions.
  * [go_BindgenFuncs](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_BindgenFuncs): Example for invokation the functions from Rust source with wasm-bindgen.
  * [go_BindgenKmeans](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_BindgenKmeans): Example for kmeans calculation with wasm-bindgen.
  * [go_BindgenWasi](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_BindgenWasi): Invoke WASM with WASI and wasm-bindgen.
* Mobilenet examples
  * [go_Mobilenet](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_Mobilenet): Basic mobilenet example with wasm-bindgen.
  * [go_MobilenetBirds](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetBirds), [go_MobilenetFood](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetFood), [go_MobilenetInsects](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetInsects), and [go_MobilenetPlants](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_MobilenetPlants): MobileNet examples for every categories.
* Tensorflow extension examples
  * [go_mtcnn](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_mtcnn): WasmEdge-tensorflow extension in WasmEdge-Go example.
