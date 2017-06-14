package definitions

import (
	"fmt"

	"github.com/mcollinge/disassemble/bit"
)

type ClassDefinition struct {
	Magic             uint32
	Minor             uint16
	Major             uint16
	ConstantPoolCount uint16
	ConstantPool      []ConstantDefinition
	Access            uint16
	Name              uint16
	Super             uint16
	InterfaceCount    uint16
	Interfaces        []uint16
	FieldCount        uint16
	Fields            []ClassMemberDefinition
	MethodCount       uint16
	Methods           []ClassMemberDefinition
	AttributeCount    uint16
	Attributes        []AttributeDefinition
}

func (c *ClassDefinition) Unpack(reader *bit.BinaryReader) error {

	reader.Uint32(&c.Magic)
	if c.Magic != 0xCAFEBABE {
		return fmt.Errorf("Magic does not equal 0xCAFEBABE, this files magic is ", c.Magic)
	}
	reader.Uint16(&c.Minor).
		Uint16(&c.Major).
		Uint16(&c.ConstantPoolCount)

	c.ConstantPool = make([]ConstantDefinition, c.ConstantPoolCount-1)
	for i := 1; i < int(c.ConstantPoolCount); i++ {
		constantDef := ConstantDefinition{}
		constantDef.Unpack(reader)
		c.ConstantPool[i-1] = constantDef
		if constantDef.Tag == LongConstant || constantDef.Tag == DoubleConstant {
			i++
			c.ConstantPool[i-1] = constantDef
		}
	}
	reader.Uint16(&c.Access).
		Uint16(&c.Name).
		Uint16(&c.Super).
		Uint16(&c.InterfaceCount)

	c.Interfaces = make([]uint16, c.InterfaceCount)
	for i := 0; i < int(c.InterfaceCount); i++ {
		reader.Uint16(&c.Interfaces[i])
	}

	reader.Uint16(&c.FieldCount)
	c.Fields = make([]ClassMemberDefinition, c.FieldCount)
	for i := 0; i < int(c.FieldCount); i++ {
		field := ClassMemberDefinition{}
		field.Unpack(reader)
		c.Fields[i] = field
	}

	reader.Uint16(&c.MethodCount)
	c.Methods = make([]ClassMemberDefinition, c.MethodCount)
	for i := 0; i < int(c.MethodCount); i++ {
		method := ClassMemberDefinition{}
		method.Unpack(reader)
		c.Methods[i] = method
	}

	c.Attributes = make([]AttributeDefinition, c.AttributeCount)
	for i := 0; i < int(c.AttributeCount); i++ {
		attr := AttributeDefinition{}
		attr.Unpack(reader)
		c.Attributes[i] = attr
	}
	return reader.Error()
}
