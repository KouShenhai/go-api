package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"go-api/util"
	"testing"
)

func Test1(t *testing.T) {
	// a0a0a0a00001210100c60007001c0001000100a40000000300000000000000000000000000000000001c00010001008a0000000100000000000000000000000000000000001c0001000100880000000300000000000000000000000000000000001c0001000100880000000400000000000000000000000000000000001c0001000100240000000000000000000000000000000000000000001c0001000100250000000000000000000000000000000000000000001c0001000100a40000000300000000000000000000000000000000
	array := toBuildReadArray(
		toReadSingle(0xA4, 3, 0, 0, 0, 0),
		toReadSingle(0x8A, 1, 0, 0, 0, 0),
		toReadSingle(0x88, 3, 0, 0, 0, 0),
		toReadSingle(0x88, 4, 0, 0, 0, 0),
		toReadSingle(0x24, 0, 0, 0, 0, 0),
		toReadSingle(0x25, 0, 0, 0, 0, 0),
		toReadSingle(0xA4, 3, 0, 0, 0, 0),
	)
	h := hex.EncodeToString(array)
	fmt.Println(h)
}

func toReadSingle(a, b, c, d, e, f int32) []byte {
	var buffer = make([]byte, 28)
	buffer[1] = 0x1C
	buffer[3] = 0x01
	buffer[5] = 0x01
	copy(buffer[6:], util.TransByte(a))
	copy(buffer[10:], util.TransByte(b))
	copy(buffer[14:], util.TransByte(c))
	copy(buffer[18:], util.TransByte(d))
	copy(buffer[22:], util.TransByte(e))
	copy(buffer[26:], util.TransByte(f))
	return buffer
}

func toBuildReadArray(bs ...[]byte) []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte{0xa0, 0xa0, 0xa0, 0xa0, 0x00, 0x01, 0x21, 0x01, 0x00, 0x1e})
	buffer.Write(util.TransByte(int32(len(bs))))
	for _, b := range bs {
		buffer.Write(b)
	}
	result := buffer.Bytes()
	length := util.TransByte(int32(len(result) - 10))
	copy(result[8:], length)
	return result
}
