package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

const x float64 = -0.743644786
const y float64 = 0.1318252536
const d float64 = 0.00029336
const maxiterations uint32 = 10000
const width uint = 1200
const height uint = 800

func RunMandelbrot(vm *wasmedge.VM, wg *sync.WaitGroup, nthreads uint32, rank uint32) {
	_, err := vm.Execute("mandelbrotThread", maxiterations, nthreads, rank, x, y, d)
	if err == nil {
		// fmt.Println("Go: Thread rank", rank, "execution succeeded")
	} else {
		fmt.Println("Go: ERROR - Thread rank", rank, "failed:", err.Error())
	}
	wg.Done()
}

func main() {
	fmt.Println("Go: Args:", os.Args)
	// Expected Args[0]: program name (./threads)
	// Expected Args[1]: wasm name (mandelbrot.so)
	// Expected Args[2] (optional): num of threads, default 4

	// Check the command line arguments
	var nthreads uint32 = 4
	if len(os.Args) <= 1 {
		fmt.Println("Go: ERROR - WASM file needed.")
		return
	}
	fmt.Println("Go: Input WASM file:", os.Args[1])
	if len(os.Args) == 3 {
		i, err := strconv.Atoi(os.Args[2])
		if err == nil {
			nthreads = uint32(i)
		} else {
			fmt.Println("Go: ERROR - Input thread failed:", err.Error())
			return
		}
	}
	fmt.Println("Go: Num of threads:", nthreads)

	var wg sync.WaitGroup
	var res []interface{}
	var err error

	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.MULTI_MEMORIES, wasmedge.THREADS)
	conf.RemoveConfig(wasmedge.REFERENCE_TYPES)

	// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)

	// Create host memory
	lim := wasmedge.NewLimitSharedWithMax(60, 60)
	memtype := wasmedge.NewMemoryType(lim)
	mem := wasmedge.NewMemory(memtype)
	memtype.Release()

	// Create host module and add the host memory
	hostmod := wasmedge.NewModule("env")
	hostmod.AddMemory("memory", mem)

	// Register the host module
	err = vm.RegisterModule(hostmod)
	if err != nil {
		fmt.Println("Go: ERROR - Registration the host module failed:", err.Error())
		return
	}

	// Instantiate the WASM module
	err = vm.LoadWasmFile(os.Args[1])
	if err != nil {
		fmt.Println("Go: ERROR - Load the WASM failed:", err.Error())
		return
	}
	err = vm.Validate()
	if err != nil {
		fmt.Println("Go: ERROR - Validate the WASM failed:", err.Error())
		return
	}
	err = vm.Instantiate()
	if err != nil {
		fmt.Println("Go: ERROR - Instantiate the WASM failed:", err.Error())
		return
	}

	// Execute "mandelbrotThread" WASM function in multi-thread
	// Note: In WasmEdge, the goroutine threads should larger or equal to the
	//       goroutine nums, or the execution will failed.
	wg.Add(int(nthreads))
	for i := 0; i < int(nthreads); i++ {
		tid := uint32(i)
		go RunMandelbrot(vm, &wg, nthreads, tid)
	}
	wg.Wait()

	// Get image offset in the memory instance
	res, err = vm.Execute("getImage")
	if err != nil {
		fmt.Println("Go: ERROR - Get result image failed:", err.Error())
		return
	}
	offset := uint(res[0].(int32))
	fmt.Println("Go: Got the result image offset:", offset)

	// Get image data from the memory instance
	var data []byte
	mem = hostmod.FindMemory("memory")
	if mem == nil {
		fmt.Println("Go: ERROR - Cannot find the memory from the host module")
		return
	}
	data, err = mem.GetData(offset, width*height*4)
	if err != nil {
		fmt.Println("Go: ERROR - Cannot get the image data:", err.Error())
		return
	}
	os.WriteFile("output-wasmedge-bin", data, 0644)

	hostmod.Release()
	vm.Release()
	conf.Release()
}
