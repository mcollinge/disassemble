package utils

import (
	"encoding/binary"
)

type BinaryReader struct {
	bytes  []byte
	endian binary.ByteOrder
	err    error
}

func NewReader(bytes []byte, endian binary.ByteOrder) *BinaryReader {
	return &BinaryReader{bytes: bytes, endian: endian, err: nil}
}

func (this *BinaryReader) Error() error {
	return this.err
}

func (this *BinaryReader) Uint8(i *uint8) *BinaryReader {
	if this.err != nil {
		return this
	}
	*i = this.bytes[:1][0]
	this.bytes = this.bytes[1:]
	return this
}

func (this *BinaryReader) Uint16(i *uint16) *BinaryReader {
	if this.err != nil {
		return this
	}
	bytes := this.bytes[:2]
	*i = this.endian.Uint16(bytes)
	this.bytes = this.bytes[2:]
	return this
}

func (this *BinaryReader) Uint32(i *uint32) *BinaryReader {
	if this.err != nil {
		return this
	}
	bytes := this.bytes[:4]
	*i = this.endian.Uint32(bytes)
	this.bytes = this.bytes[4:]
	return this
}

func (this *BinaryReader) Uint64(i *uint64) *BinaryReader {
	if this.err != nil {
		return this
	}
	bytes := this.bytes[:8]
	*i = this.endian.Uint64(bytes)
	this.bytes = this.bytes[8:]
	return this
}

func (this *BinaryReader) String(i uint16, str *string) *BinaryReader {
	if this.err != nil {
		return this
	}
	bytes := this.bytes[:int(i)]
	*str = string(bytes)
	this.bytes = this.bytes[int(i):]
	return this
}

func (this *BinaryReader) Bytes(bytes *[]byte) *BinaryReader {
	if this.err != nil {
		return this
	}
	length := len(*bytes)
	*bytes = this.bytes[:length]
	this.bytes = this.bytes[length:]
	return this
}
