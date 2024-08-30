package test

import (
	"encoding/hex"
	"fmt"
	"go-api/util"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(hex.EncodeToString(toReadSingle(0x34, 1, 1, 3, 4, 5)))
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
