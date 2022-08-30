package main

import (
	"strings"
	"unsafe"
)

func main() {}

// export greet
func greet(subject *int32) *byte {
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

	output := subjectStr.String()
	output = "Hello, " + output + "!"

	return &(([]byte)(output)[0])
}
