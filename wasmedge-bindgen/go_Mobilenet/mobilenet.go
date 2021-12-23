package main

import (
	"fmt"
	"io/ioutil"
	"os"

	host "github.com/second-state/wasmedge-bindgen/host/go"
)

func main() {
	fmt.Println("Go: Args:", os.Args)
	/// Expected Args[0]: program name (./mobilenet)
	/// Expected Args[1]: wasm or wasm-so file (rust_mobilenet_lib_bg.wasm)
	/// Expected Args[2]: input image name (solvay.jpg)

	/// Set Tensorflow not to print debug info
	os.Setenv("TF_CPP_MIN_LOG_LEVEL", "3")
	os.Setenv("TF_CPP_MIN_VLOG_LEVEL", "3")

	wh, err := host.NewHost(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wh.Release()
	
	wh.InstallTensorflowExt()

	img, _ := ioutil.ReadFile(os.Args[2])
	if res, err := wh.Run("infer", img); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(res))
	}
}