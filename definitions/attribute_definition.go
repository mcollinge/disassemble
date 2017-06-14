package definitions

import "github.com/mcollinge/disassemble/bit"

type AttributeDefinition struct {
	NameIndex       uint16
	AttributeLength uint32
	Info            []byte
}

func (c *AttributeDefinition) Unpack(reader *bit.BinaryReader) error {
	reader.Uint16(&c.NameIndex).
		Uint32(&c.AttributeLength)

	c.Info = make([]byte, int(c.AttributeLength))
	reader.Bytes(&c.Info)
	return reader.Error()
}
