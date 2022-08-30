package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func main() {
	// Expected Args[0]: program name (./bindgen_kmeans)
	// Expected Args[1]: wasm or wasm-so file (rust_bindgen_kmeans_lib_bg.wasm))

	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)

	// Init WASI
	var wasi = vm.GetImportModule(wasmedge.WASI)
	wasi.InitWasi(
		os.Args[1:],     // The args
		os.Environ(),    // The envs
		[]string{".:."}, // The mapping preopens
	)

	// Instantiate wasm
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()
	vm.Instantiate()

	// Run bindgen functions
	var res interface{}
	var err error
	var csv []byte
	// fit: array, i32, i32 -> array
	csv, err = ioutil.ReadFile("birch3.data.csv")
	res, err = vm.ExecuteBindgen("fit", wasmedge.Bindgen_return_array, csv, int32(2), int32(100))
	if err == nil {
		fmt.Println("Run bindgen -- fit (birch3 cluster centers):", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- fit (birch3 cluster centers) FAILED")
	}
	// fit: array, i32, i32 -> array
	csv, err = ioutil.ReadFile("iris.data.csv")
	res, err = vm.ExecuteBindgen("fit", wasmedge.Bindgen_return_array, csv, int32(2), int32(3))
	if err == nil {
		fmt.Println("Run bindgen -- fit (iris cluster centers):", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- fit (iris cluster centers) FAILED")
	}
	// fit: array, i32, i32 -> array
	csv, err = ioutil.ReadFile("s1.data.csv")
	res, err = vm.ExecuteBindgen("fit", wasmedge.Bindgen_return_array, csv, int32(2), int32(15))
	if err == nil {
		fmt.Println("Run bindgen -- fit (s1 cluster centers):", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- fit (s1 cluster centers) FAILED")
	}
	// fit: array, i32, i32 -> array
	csv, err = ioutil.ReadFile("dim128.data.csv")
	res, err = vm.ExecuteBindgen("fit", wasmedge.Bindgen_return_array, csv, int32(128), int32(16))
	if err == nil {
		fmt.Println("Run bindgen -- fit (dim128 cluster centers):", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- fit (dim128 cluster centers) FAILED")
	}

	vm.Release()
	conf.Release()
}
