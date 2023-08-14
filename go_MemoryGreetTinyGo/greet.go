package main

import (
	"strings"
	"unsafe"
)

func main() {}

//export greet
func greet(subject *int32) *int32 {
	nth := 0
	var subjectStr strings.Builder
	pointer := uintptr(unsafe.Pointer(subject))
	for {
		s := *(*int32)(unsafe.Pointer(pointer + uintptr(nth)))
		if s == 0 {
			break
		}

		subjectStr.WriteByte(byte(s))
		nth++
	}

	output := []byte("Hello, " + subjectStr.String() + "!")

	r := make([]int32, 2)
	r[0] = int32(uintptr(unsafe.Pointer(&(output[0]))))
	r[1] = int32(len(output))

	return &r[0]
}
