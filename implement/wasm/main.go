package main

import (
	"encoding/hex"
	"unsafe"
)

func convert(s string) string {
	return hex.EncodeToString([]byte(s))
}

func main() {}

//export convert
func _convert(ptr, size uint32) (ptrSize uint64) {
	s := ptrToString(ptr, size)
	result := convert(s)

	p, size := stringToPtr(result)
	return (uint64(p) << uint64(32)) | uint64(size)
}

func ptrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func stringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}
