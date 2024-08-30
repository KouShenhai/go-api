package util

import "encoding/binary"

func TransByte(val int32) []byte {
	// 创建一个缓冲区
	buffer := make([]byte, 2)
	// 将整数转换为字节数组并写入缓冲区
	binary.BigEndian.PutUint16(buffer, uint16(val))
	// 打印缓冲区内容
	return buffer
}
