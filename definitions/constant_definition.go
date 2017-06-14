package definitions

import (
	"github.com/mcollinge/disassemble/bit"
)

const (
	ClassConstant              uint8 = 7
	FieldRefConstant           uint8 = 9
	MethodRefConstant          uint8 = 10
	InterfaceMethodRefConstant uint8 = 11
	StringConstant             uint8 = 8
	IntegerConstant            uint8 = 3
	FloatContstant             uint8 = 4
	LongConstant               uint8 = 5
	DoubleConstant             uint8 = 6
	NameAndTypeConstant        uint8 = 12
	Utf8Constant               uint8 = 1
	MethodHandleConstant       uint8 = 15
	MethodTypeConstant         uint8 = 16
	InvokeDynmaicConstant      uint8 = 18
)

type ConstantDefinition struct {
	Tag  uint8
	Data []byte
}

func (c *ConstantDefinition) Unpack(reader *bit.BinaryReader) error {
	reader.Uint8(&c.Tag)
	switch c.Tag {
	case FieldRefConstant, MethodRefConstant, InterfaceMethodRefConstant,
		IntegerConstant, FloatContstant, NameAndTypeConstant, InvokeDynmaicConstant:
		c.Data = make([]byte, 4)
		reader.Bytes(&c.Data)

	case LongConstant, DoubleConstant:
		c.Data = make([]byte, 8)
		reader.Bytes(&c.Data)

	case Utf8Constant:
		var length uint16
		reader.Uint16(&length)
		c.Data = make([]byte, int(length))
		reader.Bytes(&c.Data)

	case MethodHandleConstant:
		c.Data = make([]byte, 3)
		reader.Bytes(&c.Data)

	case ClassConstant, MethodTypeConstant, StringConstant:
		c.Data = make([]byte, 2)
		reader.Bytes(&c.Data)
	}
	return nil
}
