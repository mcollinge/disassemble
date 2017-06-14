package definitions

import "github.com/mcollinge/disassemble/bit"

type ClassMemberDefinition struct {
	Access          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributeCount  uint16
	Attributes      []AttributeDefinition
}

func (c *ClassMemberDefinition) Unpack(reader *bit.BinaryReader) error {
	reader.Uint16(&c.Access).
		Uint16(&c.NameIndex).
		Uint16(&c.DescriptorIndex).
		Uint16(&c.AttributeCount)

	c.Attributes = make([]AttributeDefinition, c.AttributeCount)
	for i := 0; i < int(c.AttributeCount); i++ {
		attr := AttributeDefinition{}
		attr.Unpack(reader)
		c.Attributes[i] = attr
	}
	return reader.Error()
}
