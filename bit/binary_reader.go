package bit

import (
	"encoding/binary"
	"io"
	"io/ioutil"
)

type BinaryReader struct {
	bytes  []byte
	index  int
	endian binary.ByteOrder
	err    error
}

func NewReader(reader io.Reader, endian binary.ByteOrder) (*BinaryReader, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return &BinaryReader{bytes: bytes, index: 0, endian: endian, err: nil}, nil
}

func (br *BinaryReader) Error() error {
	return br.err
}

func (br *BinaryReader) Uint8(i *uint8) *BinaryReader {
	if br.err != nil {
		return br
	}
	*i = br.bytes[br.index]
	br.index += 1
	return br
}

func (br *BinaryReader) Uint16(i *uint16) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := br.bytes[br.index : br.index+2]
	br.index += 2
	*i = br.endian.Uint16(bytes)
	return br
}

func (br *BinaryReader) Uint32(i *uint32) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := br.bytes[br.index : br.index+4]
	br.index += 4
	*i = br.endian.Uint32(bytes)
	return br
}

func (br *BinaryReader) Uint64(i *uint64) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := br.bytes[br.index : br.index+8]
	br.index += 8
	*i = br.endian.Uint64(bytes)
	return br
}

func (br *BinaryReader) String(i uint16, str *string) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := br.bytes[br.index : br.index+int(i)]
	br.index += int(i)
	*str = string(bytes)
	return br
}

func (br *BinaryReader) Bytes(bytes *[]byte) *BinaryReader {
	if br.err != nil {
		return br
	}
	length := len(*bytes)
	*bytes = br.bytes[br.index : br.index+length]
	br.index += length
	return br
}
