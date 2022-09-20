package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
	bindgen "github.com/second-state/wasmedge-bindgen/host/go"
)

func main() {
	// Expected Args[0]: program name (./bindgen_wasi)
	// Expected Args[1]: wasm or wasm-so file (rust_bindgen_wasi_lib_bg.wasm))

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

	vm.LoadWasmFile(os.Args[1])
	vm.Validate()
	// Instantiate the bindgen and vm
	bg := bindgen.New(vm)
	bg.Instantiate()

	// Run bindgen functions
	var res []interface{}
	var err error
	// get_random_i32: {} -> i32
	res, _, err = bg.Execute("get_random_i32")
	if err == nil {
		fmt.Println("Run bindgen -- get_random_i32:", res[0].(int32))
	} else {
		fmt.Println("Run bindgen -- get_random_i32 FAILED")
	}
	// get_random_bytes: {} -> array
	res, _, err = bg.Execute("get_random_bytes")
	if err == nil {
		fmt.Println("Run bindgen -- get_random_bytes:", res[0].([]byte))
	} else {
		fmt.Println("Run bindgen -- get_random_bytes FAILED")
	}
	// echo: array -> array
	res, _, err = bg.Execute("echo", "hello!!!!")
	if err == nil {
		fmt.Println("Run bindgen -- echo:", res[0].(string))
	} else {
		fmt.Println("Run bindgen -- echo FAILED")
	}
	// print_env: {} -> {}
	res, _, err = bg.Execute("print_env")
	if err == nil {
		fmt.Println("Run bindgen -- print_env")
	} else {
		fmt.Println("Run bindgen -- print_env FAILED")
	}
	// create_file: array, array -> {}
	res, _, err = bg.Execute("create_file", "test.txt", "TEST MESSAGES----!@#@%@%$#!@#")
	if err == nil {
		fmt.Println("Run bindgen -- create_file: test.txt")
	} else {
		fmt.Println("Run bindgen -- create_file FAILED")
	}

	// read_file: array -> array
	res, _, err = bg.Execute("read_file", "test.txt")
	if err == nil {
		fmt.Println("Run bindgen -- read_file:", res[0].(string))
	} else {
		fmt.Println("Run bindgen -- read_file FAILED")
	}

	// del_file: array -> {}
	res, _, err = bg.Execute("del_file", "test.txt")
	if err == nil {
		fmt.Println("Run bindgen -- del_file: test.txt")
	} else {
		fmt.Println("Run bindgen -- del_file FAILED")
	}

	bg.Release()
	conf.Release()
}
