package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

// The real worker functions
func real_add(a int32, b int32) int32 {
	fmt.Println("Go: Entering go function real_add")
	c := a + b
	fmt.Println("Go: Leaving go function real_add")
	return c
}
func real_mul(a int32, b int32) int32 {
	fmt.Println("Go: Entering go function real_mul")
	c := a * b
	fmt.Println("Go: Leaving go function real_mul")
	return c
}
func real_square(a int32) int32 {
	fmt.Println("Go: Entering go function real_square")
	c := a * a
	fmt.Println("Go: Leaving go function real_square")
	return c
}

// Host functions
func host_add(data interface{}, callframe *wasmedge.CallingFrame, params []interface{}) ([]interface{}, wasmedge.Result) {
	// add: externref, i32, i32 -> i32
	// call the real add function in externref
	fmt.Println("Go: Entering go host function host_add")

	// Get the externref
	externref := params[0].(wasmedge.ExternRef)

	// Get the interface{} from externref
	realref := externref.GetRef()

	// Cast to the function
	realfunc := realref.(func(int32, int32) int32)

	// Call function
	res := realfunc(params[1].(int32), params[2].(int32))

	// Set the returns
	returns := make([]interface{}, 1)
	returns[0] = res

	// Return
	fmt.Println("Go: Leaving go host function host_add")
	return returns, wasmedge.Result_Success
}

func host_mul(data interface{}, callframe *wasmedge.CallingFrame, params []interface{}) ([]interface{}, wasmedge.Result) {
	// mul: externref, i32, i32 -> i32
	// call the real mul function in externref
	fmt.Println("Go: Entering go host function host_mul")

	// Get the externref
	externref := params[0].(wasmedge.ExternRef)

	// Get the interface{} from externref
	realref := externref.GetRef()

	// Cast to the function
	realfunc := realref.(func(int32, int32) int32)

	// Call function
	res := realfunc(params[1].(int32), params[2].(int32))

	// Set the returns
	returns := make([]interface{}, 1)
	returns[0] = res

	// Return
	fmt.Println("Go: Leaving go host function host_mul")
	return returns, wasmedge.Result_Success
}

func host_square(data interface{}, callframe *wasmedge.CallingFrame, params []interface{}) ([]interface{}, wasmedge.Result) {
	// square: externref, i32 -> i32
	// call the real square function in externref
	fmt.Println("Go: Entering go host function host_square")

	// Get the externref
	externref := params[0].(wasmedge.ExternRef)

	// Get the interface{} from externref
	realref := externref.GetRef()

	// Cast to the function
	realfunc := realref.(func(int32) int32)

	// Call function
	res := realfunc(params[1].(int32))

	// Set the returns
	returns := make([]interface{}, 1)
	returns[0] = res

	// Return
	fmt.Println("Go: Leaving go host function host_square")
	return returns, wasmedge.Result_Success
}

func main() {
	fmt.Println("Go: Args:", os.Args)
	// Expected Args[0]: program name (./externref)
	// Expected Args[1]: wasm file (funcs.wasm)

	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Create VM
	var vm = wasmedge.NewVM()

	// Create module instance
	var obj = wasmedge.NewModule("extern_module")

	// Add host functions into the module instance
	var type1 = wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_ExternRef,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
		}, []wasmedge.ValType{
			wasmedge.ValType_I32,
		})
	var type2 = wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_ExternRef,
			wasmedge.ValType_I32,
		}, []wasmedge.ValType{
			wasmedge.ValType_I32,
		})
	var func_add = wasmedge.NewFunction(type1, host_add, nil, 0)
	var func_mul = wasmedge.NewFunction(type1, host_mul, nil, 0)
	var func_square = wasmedge.NewFunction(type2, host_square, nil, 0)
	obj.AddFunction("add", func_add)
	obj.AddFunction("mul", func_mul)
	obj.AddFunction("square", func_square)

	// Register module instance
	vm.RegisterModule(obj)

	// Instantiate wasm
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()
	vm.Instantiate()

	// Run
	var ref_add = wasmedge.NewExternRef(real_add)
	var ref_mul = wasmedge.NewExternRef(real_mul)
	var ref_square = wasmedge.NewExternRef(real_square)
	var res []interface{}
	var err error
	res, err = vm.Execute("call_add", ref_add, int32(1234), int32(5678))
	if err == nil {
		fmt.Println("Run call_add: 1234 + 5678 =", res[0].(int32))
	} else {
		fmt.Println("Run call_add FAILED")
	}
	res, err = vm.Execute("call_mul", ref_mul, int32(4827), int32(-31519))
	if err == nil {
		fmt.Println("Run call_mul: 4827 * (-31519) =", res[0].(int32))
	} else {
		fmt.Println("Run call_mul FAILED")
	}
	res, err = vm.Execute("call_square", ref_square, int32(1024))
	if err == nil {
		fmt.Println("Run call_square: 1024^2 =", res[0].(int32))
	} else {
		fmt.Println("Run call_square FAILED")
	}
	res, err = vm.Execute("call_add_square", ref_add, ref_square, int32(761), int32(195))
	if err == nil {
		fmt.Println("Run call_square: (761 + 195)^2 =", res[0].(int32))
	} else {
		fmt.Println("Run call_square FAILED")
	}

	ref_add.Release()
	ref_mul.Release()
	ref_square.Release()
	vm.Release()
	obj.Release()
}
