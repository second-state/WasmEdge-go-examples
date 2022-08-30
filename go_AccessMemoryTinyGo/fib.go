package main

import (
	"fmt"
	"unsafe"
)

func main() {
	println("in main")
	n := int32(10)
	arr := make([]int32, n)
	arrP := &arr[0]
	fmt.Printf("call fibArray(%d, %p) = %d\n", n, arrP, fibArray(n, arrP))
	fmt.Printf("call fibArrayReturnMemory(%d) return %p\n", n, fibArrayReturnMemory(n))
}

// export fibArray
func fibArray(n int32, p *int32) int32 {
	arr := unsafe.Slice(p, n)
	for i := int32(0); i < n; i++ {
		switch {
		case i < 2:
			arr[i] = i
		default:
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr[n-1]
}

// export fibArrayReturnMemory
func fibArrayReturnMemory(n int32) *int32 {
	arr := make([]int32, n)
	for i := int32(0); i < n; i++ {
		switch {
		case i < 2:
			arr[i] = i
		default:
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return &arr[0]
}
