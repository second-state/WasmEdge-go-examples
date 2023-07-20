package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
	bindgen "github.com/second-state/wasmedge-bindgen/host/go"
)

func main() {
	fmt.Println("Go: Args:", os.Args)
	// Expected Args[0]: program name (./mobilenet)
	// Expected Args[1]: wasm or wasm-so file (rust_mobilenet_lib_bg.wasm)
	// Expected Args[2]: input image name (solvay.jpg)

	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Set Tensorflow not to print debug info
	os.Setenv("TF_CPP_MIN_LOG_LEVEL", "3")
	os.Setenv("TF_CPP_MIN_VLOG_LEVEL", "3")

	// Load WasmEdge-image and WasmEdge-tensorflow from default path
	wasmedge.LoadPluginDefaultPaths()

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

	// Load and validate the wasm
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()

	// Instantiate the bindgen and vm
	bg := bindgen.New(vm)
	bg.Instantiate()

	img, _ := ioutil.ReadFile(os.Args[2])
	res, _, err := bg.Execute("infer", img)
	if err == nil {
		fmt.Println("GO: Run bindgen -- infer:", res[0].(string))
	} else {
		fmt.Println("GO: Run bindgen -- infer FAILED")
	}

	vm.Release()
	conf.Release()
}
