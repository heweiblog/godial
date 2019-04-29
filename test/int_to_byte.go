package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func cal_mask(val int32) uint32 {
	var i int32 = 0
	var res uint32 = 1
	for i = 0; i < val; i++ {
		res *= 2
	}
	res -= 1
	return ^res
}

func IntToBytes(n uint32) []byte {
	x := uint32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func main() {
	var a uint32 = cal_mask(8)
	fmt.Println(IntToBytes(a))
}
