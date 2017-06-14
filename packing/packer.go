package packing

import (
	"encoding/binary"
	"io"

	"github.com/mcollinge/disassemble/bit"
	"github.com/mcollinge/disassemble/definitions"
)

func Unpack(reader io.Reader, def definitions.Definition) error {
	binaryReader, err := bit.NewReader(reader, binary.BigEndian)
	if err != nil {
		return err
	}
	def.Unpack(binaryReader)
	return binaryReader.Error()
}
