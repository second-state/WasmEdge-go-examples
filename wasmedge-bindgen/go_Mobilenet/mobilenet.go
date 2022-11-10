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
	// Expected Args[1]: wasm file (rust_mobilenet_lib.wasm)
	// Expected Args[2]: input image name (grace_hopper.jpg)

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

	// Register WasmEdge-tensorflow
	var tfobj = wasmedge.NewTensorflowModule()
	var tfliteobj = wasmedge.NewTensorflowLiteModule()
	vm.RegisterModule(tfobj)
	vm.RegisterModule(tfliteobj)

	// Load and validate the wasm
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()

	// Instantiate the bindgen and vm
	bg := bindgen.New(vm)
	bg.Instantiate()

	img, _ := ioutil.ReadFile(os.Args[2])
	if res, _, err := bg.Execute("infer", img); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res[0].(string))
	}

	bg.Release()
	vm.Release()
	conf.Release()
	tfobj.Release()
	tfliteobj.Release()
}
