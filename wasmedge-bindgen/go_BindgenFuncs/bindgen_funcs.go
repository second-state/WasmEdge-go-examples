package main

import (
	"fmt"
	"os"
	"strconv"

	host "github.com/second-state/wasmedge-bindgen/host/go"
)

func main() {
	/// Expected Args[0]: program name (./bindgen_funcs)
	/// Expected Args[1]: wasm or wasm-so file (rust_bindgen_funcs_lib_bg.wasm))

	wh, err := host.NewHost(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wh.Release()

	/// create_line: string, string, string -> string (inputs are JSON stringified)	
	res, err := wh.Run("create_line", "{\"x\":2.5,\"y\":7.8}", "{\"x\":2.5,\"y\":5.8}", "A thin red line")
	if err == nil {
		fmt.Println("Run bindgen -- create_line:", string(res))
	} else {
		fmt.Println("Run bindgen -- create_line FAILED", err)
	}

	/// say: string -> string
	res, err = wh.Run("say", "bindgen funcs test")
	if err == nil {
		fmt.Println("Run bindgen -- say:", string(res))
	} else {
		fmt.Println("Run bindgen -- say FAILED")
	}

	/// obfusticate: string -> string
	res, err = wh.Run("obfusticate", "A quick brown fox jumps over the lazy dog")
	if err == nil {
		fmt.Println("Run bindgen -- obfusticate:", string(res))
	} else {
		fmt.Println("Run bindgen -- obfusticate FAILED")
	}

	/// lowest_common_multiple: i32, i32 -> i32
	res, err = wh.Run("lowest_common_multiple", int32(123), int32(2))
	if err == nil {
		num, _ := strconv.ParseInt(string(res), 10, 32)
		fmt.Println("Run bindgen -- lowest_common_multiple:", num)
	} else {
		fmt.Println("Run bindgen -- lowest_common_multiple FAILED")
	}

	/// sha3_digest: array -> array
	res, err = wh.Run("sha3_digest", []byte("This is an important message"))
	if err == nil {
		fmt.Println("Run bindgen -- sha3_digest:", res)
	} else {
		fmt.Println("Run bindgen -- sha3_digest FAILED")
	}

	/// keccak_digest: array -> array
	res, err = wh.Run("keccak_digest", []byte("This is an important message"))
	if err == nil {
		fmt.Println("Run bindgen -- keccak_digest:", res)
	} else {
		fmt.Println("Run bindgen -- keccak_digest FAILED")
	}
}
