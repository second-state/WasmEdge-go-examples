package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func main() {
	/// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	/// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	/// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)

	/// Init WASI
	var wasi = vm.GetImportModule(wasmedge.WASI)
	wasi.InitWasi(
		os.Args[1:],     /// The args
		os.Environ(),    /// The envs
		[]string{".:."}, /// The mapping preopens
	)

	/// Run WASM file
	vm.RunWasmFile(os.Args[1], "_start")

	exitcode := wasi.WasiGetExitCode()
	if exitcode != 0 {
		fmt.Println("Go: Run WASM failed, exit code:", exitcode)
	}

	vm.Release()
	conf.Release()
}
