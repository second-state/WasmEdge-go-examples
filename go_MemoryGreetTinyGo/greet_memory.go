package main

import (
	"fmt"
	"os"
	"encoding/binary"

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
	allocateResult, _ := vm.Execute("malloc", int32(lengthOfSubject+1))
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

	memData, _ = mem.GetData(uint(outputPointer), 8)
	resultPointer := binary.LittleEndian.Uint32(memData[:4])
	resultLength := binary.LittleEndian.Uint32(memData[4:])

	// Read the result of the `greet` function.
	memData, _ = mem.GetData(uint(resultPointer), uint(resultLength))
	fmt.Println(string(memData))

	// Deallocate the subject, and the output.
	vm.Execute("free", inputPointer)

	vm.Release()
	conf.Release()
}
