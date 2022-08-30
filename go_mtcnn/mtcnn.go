package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func main() {
	fmt.Println("Go: Args:", os.Args)
	// Expected Args[0]: program name (./mtcnn)
	// Expected Args[1]: wasm or wasm-so file (rust_mtcnn.wasm)
	// Expected Args[2]: model file (mtcnn.pb)
	// Expected Args[3]: input image name (solvay.jpg)
	// Expected Args[4]: output image name

	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Set Tensorflow not to print debug info
	os.Setenv("TF_CPP_MIN_LOG_LEVEL", "3")
	os.Setenv("TF_CPP_MIN_VLOG_LEVEL", "3")

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

	// Register WasmEdge-image and WasmEdge-tensorflow
	var tfobj = wasmedge.NewTensorflowModule()
	var tfliteobj = wasmedge.NewTensorflowLiteModule()
	vm.RegisterModule(tfobj)
	vm.RegisterModule(tfliteobj)

	// Run WASM file
	vm.RunWasmFile(os.Args[1], "_start")

	exitcode := wasi.WasiGetExitCode()
	if exitcode != 0 {
		fmt.Println("Go: Run WASM failed, exit code:", exitcode)
	}

	vm.Release()
	conf.Release()
	tfobj.Release()
	tfliteobj.Release()
}
