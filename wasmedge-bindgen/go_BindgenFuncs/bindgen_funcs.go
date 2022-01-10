package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
	bindgen "github.com/second-state/wasmedge-bindgen/host/go"
)

func main() {
	/// Expected Args[0]: program name (./bindgen_funcs)
	/// Expected Args[1]: wasm file (rust_bindgen_funcs_lib.wasm))
	
	/// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	/// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	/// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)

	/// Init WASI
	var wasi = vm.GetImportObject(wasmedge.WASI)
	wasi.InitWasi(
		os.Args[1:],     /// The args
		os.Environ(),    /// The envs
		[]string{".:."}, /// The mapping preopens
	)

	/// Load and validate the wasm
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()

	// Instantiate the bindgen and vm
	bg := bindgen.Instantiate(vm)

	/// create_line: string, string, string -> string (inputs are JSON stringified)	
	res, err := bg.Execute("create_line", "{\"x\":2.5,\"y\":7.8}", "{\"x\":2.5,\"y\":5.8}", "A thin red line")
	if err == nil {
		fmt.Println("Run bindgen -- create_line:", res[0].(string))
	} else {
		fmt.Println("Run bindgen -- create_line FAILED", err)
	}

	/// say: string -> string
	res, err = bg.Execute("say", "bindgen funcs test")
	if err == nil {
		fmt.Println("Run bindgen -- say:", res[0].(string))
	} else {
		fmt.Println("Run bindgen -- say FAILED")
	}

	/// obfusticate: string -> string
	res, err = bg.Execute("obfusticate", "A quick brown fox jumps over the lazy dog")
	if err == nil {
		fmt.Println("Run bindgen -- obfusticate:", res[0].(string))
	} else {
		fmt.Println("Run bindgen -- obfusticate FAILED")
	}

	/// lowest_common_multiple: i32, i32 -> i32
	res, err = bg.Execute("lowest_common_multiple", int32(123), int32(2))
	if err == nil {
		fmt.Println("Run bindgen -- lowest_common_multiple:", res[0].(int32))
	} else {
		fmt.Println("Run bindgen -- lowest_common_multiple FAILED")
	}

	/// sha3_digest: array -> array
	res, err = bg.Execute("sha3_digest", []byte("This is an important message"))
	if err == nil {
		fmt.Println("Run bindgen -- sha3_digest:", res[0].([]byte))
	} else {
		fmt.Println("Run bindgen -- sha3_digest FAILED")
	}

	/// keccak_digest: array -> array
	res, err = bg.Execute("keccak_digest", []byte("This is an important message"))
	if err == nil {
		fmt.Println("Run bindgen -- keccak_digest:", res[0].([]byte))
	} else {
		fmt.Println("Run bindgen -- keccak_digest FAILED")
	}

	bg.Release()
	vm.Release()
	conf.Release()
}
