package main

import (
	"fmt"
	"os"
	"strings"

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

	subject := "WasmEdge"
	lengthOfSubject := len(subject)

	// Allocate memory for the subject, and get a pointer to it.
	// Include a byte for the NULL terminator we add below.
	allocateResult, _ := vm.Execute("allocate", int32(lengthOfSubject+1))
	inputPointer := allocateResult[0].(int32)

	// Write the subject into the memory.
	mod := vm.GetActiveModule()
	mem := mod.FindMemory("memory")
	memData, _ := mem.GetData(uint(inputPointer), uint(lengthOfSubject+1))
	copy(memData, subject)

	// C-string terminates by NULL.
	memData[lengthOfSubject] = 0

	// Run the `greet` function. Given the pointer to the subject.
	greetResult, _ := vm.Execute("greet", inputPointer)
	outputPointer := greetResult[0].(int32)

	pageSize := mem.GetPageSize()
	// Read the result of the `greet` function.
	memData, _ = mem.GetData(uint(0), uint(pageSize*65536))
	nth := 0
	var output strings.Builder

	for {
		if memData[int(outputPointer)+nth] == 0 {
			break
		}

		output.WriteByte(memData[int(outputPointer)+nth])
		nth++
	}

	lengthOfOutput := nth

	fmt.Println(output.String())

	// Deallocate the subject, and the output.
	vm.Execute("deallocate", inputPointer, int32(lengthOfSubject+1))
	vm.Execute("deallocate", outputPointer, int32(lengthOfOutput+1))

	vm.Release()
	conf.Release()
}
