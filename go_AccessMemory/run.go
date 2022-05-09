package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func main() {
	wasmedge.SetLogErrorLevel()
	conf := wasmedge.NewConfigure(wasmedge.WASI)
	vm := wasmedge.NewVMWithConfig(conf)

	wasi := vm.GetImportModule(wasmedge.WASI)
	wasi.InitWasi(
		os.Args[1:],
		os.Environ(),
		[]string{".:."},
	)

	err := vm.LoadWasmFile(os.Args[1])
	if err != nil {
		fmt.Println("failed to load wasm")
	}
	vm.Validate()
	vm.Instantiate()

	n := int32(10)

	p, err := vm.Execute("allocate", 4*n)
	if err != nil {
		fmt.Println("allocate failed:", err)
	}

	fib, err := vm.Execute("fib_array", n, p[0])
	if err != nil {
		fmt.Println("fib_rray failed:", err)
	} else {
		fmt.Println("fib_array() returned:", fib[0])
		fmt.Printf("fib_array memory at: %p\n", unsafe.Pointer((uintptr)(p[0].(int32))))
		mod := vm.GetActiveModule()
		mem := mod.FindMemory("memory")
		if mem != nil {
			// int32 occupies 4 bytes
			fibArray, err := mem.GetData(uint(p[0].(int32)), uint(n*4))
			if err == nil && fibArray != nil {
				fmt.Println("fibArray:", fibArray)
			}
		}
	}

	fibP, err := vm.Execute("fib_array_return_memory", n)
	if err != nil {
		fmt.Println("fib_array_return_memory failed:", err)
	} else {
		fmt.Printf("fib_array_return_memory memory at: %p\n", unsafe.Pointer((uintptr)(fibP[0].(int32))))
		mod := vm.GetActiveModule()
		mem := mod.FindMemory("memory")
		if mem != nil {
			// int32 occupies 4 bytes
			fibArrayReturnMemory, err := mem.GetData(uint(fibP[0].(int32)), uint(n*4))
			if err == nil && fibArrayReturnMemory != nil {
				fmt.Println("fibArrayReturnMemory:", fibArrayReturnMemory)
			}
		}
	}

	_, err = vm.Execute("deallocate", p[0].(int32), 4*n)
	if err != nil {
		fmt.Println("free failed:", err)
	}

	exitcode := wasi.WasiGetExitCode()
	if exitcode != 0 {
		fmt.Println("Go: Running wasm failed, exit code:", exitcode)
	}

	vm.Release()
	conf.Release()
}
