package bit

import (
	"bufio"
	"encoding/binary"
	"io"
)

type BinaryReader struct {
	reader *bufio.Reader
	endian binary.ByteOrder
	err    error
}

func NewReader(reader io.Reader, endian binary.ByteOrder) *BinaryReader {
	return &BinaryReader{reader: bufio.NewReader(reader), endian: endian, err: nil}
}

func (br *BinaryReader) Finish() error {
	return br.err
}

func (br *BinaryReader) Byte(b *byte) *BinaryReader {
	if br.err != nil {
		return br
	}
	*b, br.err = br.reader.ReadByte()
	return br
}

func (br *BinaryReader) Uint16(i *uint16) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := make([]byte, 2)
	_, err := br.reader.Read(bytes)
	if err != nil {
		br.err = err
	}
	*i = br.endian.Uint16(bytes)
	return br
}

func (br *BinaryReader) Uint32(i *uint32) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := make([]byte, 4)
	read, err := br.reader.Read(bytes)
	println(read)
	if err != nil {
		println(err)
		br.err = err
	}
	*i = br.endian.Uint32(bytes)
	return br
}

func (br *BinaryReader) Uint64(i *uint64) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := make([]byte, 4)
	_, err := br.reader.ReadBytes(8)
	if err != nil {
		br.err = err
	}
	*i = br.endian.Uint64(bytes)
	return br
}

func (br *BinaryReader) String(i uint16, str *string) *BinaryReader {
	if br.err != nil {
		return br
	}
	bytes := make([]byte, 4)
	_, err := br.reader.Read(bytes)
	if err != nil {
		br.err = err
	}
	*str = string(bytes)
	return br
}
