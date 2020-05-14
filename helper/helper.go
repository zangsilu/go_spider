package helper

import (
	"bytes"
	"encoding/binary"
)

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	_ = binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	_ = binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}
