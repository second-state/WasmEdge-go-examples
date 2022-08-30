package main

import (
	"fmt"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func HostPrint(data interface{}, callframe *wasmedge.CallingFrame, param []interface{}) ([]interface{}, wasmedge.Result) {
	// param[0]: external reference
	ref := param[0].(wasmedge.ExternRef)
	value := ref.GetRef().(*int32)
	// param[1]: result of fibonacci
	fmt.Println(" [HostFunction] external value: ", *value, " , fibonacci number: ", param[1].(int32))
	return []interface{}{}, wasmedge.Result_Success
}

func ListInsts(name *string, mod *wasmedge.Module) {
	if name == nil {
		fmt.Println(" --- Exported instances of the anonymous module")
	} else {
		fmt.Println(" --- Exported instances of the module", *name)
	}
	nf := mod.ListFunction()
	fmt.Println("     --- Functions (", len(nf), ") : ", nf)
	nt := mod.ListTable()
	fmt.Println("     --- Tables    (", len(nt), ") : ", nt)
	nm := mod.ListMemory()
	fmt.Println("     --- Memories  (", len(nm), ") : ", nm)
	ng := mod.ListGlobal()
	fmt.Println("     --- Globals   (", len(ng), ") : ", ng)
}

func main() {
	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	// Create store
	var store = wasmedge.NewStore()

	// Create VM by configure and external store
	var vm = wasmedge.NewVMWithConfigAndStore(conf, store)

	// Create module instance
	var impobj = wasmedge.NewModule("host")

	// Create host function
	var hostftype = wasmedge.NewFunctionType(
		[]wasmedge.ValType{wasmedge.ValType_ExternRef, wasmedge.ValType_I32},
		[]wasmedge.ValType{})
	var hostprint = wasmedge.NewFunction(hostftype, HostPrint, nil, 0)
	hostftype.Release()

	// Add host functions into module instance
	impobj.AddFunction("print_val_and_res", hostprint)

	// Register import module as module name "host"
	vm.RegisterModule(impobj)

	// Register fibonacci wasm as module name "wasm"
	vm.RegisterWasmFile("wasm", "fibonacci.wasm")

	// Instantiate wasm
	vm.LoadWasmFile("test.wasm")
	vm.Validate()
	vm.Instantiate()

	// -----------logging-------------
	modlist := store.ListModule()
	fmt.Println("registered modules: ", modlist)
	ListInsts(nil, vm.GetActiveModule())
	for _, name := range modlist {
		ListInsts(&name, store.FindModule(name))
	}
	// -----------logging-------------

	// Create external reference
	var value int32 = 123456
	refval := wasmedge.NewExternRef(&value)

	// Run print external value 123456 and fib[20]
	fmt.Println(" ### Running print_val_and_fib with fib[", 20, "] ...")
	var _, err = vm.Execute("print_val_and_fib", refval, uint32(20))
	if err != nil {
		fmt.Println(" !!! Error: ", err.Error())
	}

	// Run print external value 876543210 and fib[21]
	value = 876543210
	fmt.Println(" ### Running print_val_and_fib with fib[", 21, "] ...")
	_, err = vm.Execute("print_val_and_fib", refval, uint32(21))
	if err != nil {
		fmt.Println(" !!! Error: ", err.Error())
	}

	// Run fib[22] directly
	fmt.Println(" ### Running wasm::fib[", 22, "] ...")
	ret, err2 := vm.ExecuteRegistered("wasm", "fib", uint32(22))
	if err2 != nil {
		fmt.Println(" !!! Error: ", err.Error())
	} else if ret != nil {
		for _, val := range ret {
			fmt.Println(" Return value: ", val)
		}
	}

	refval.Release()
	vm.Release()
	conf.Release()
	store.Release()
	impobj.Release()
}
