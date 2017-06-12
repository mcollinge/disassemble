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
}

func (c *ClassDefinition) Unpack(reader *bit.BinaryReader) error {

	reader.Uint32(&c.Magic)
	if c.Magic != 0xCAFEBABE {
		return fmt.Errorf("Magic does not equal 0xCAFEBABE, this files magic is ", c.Magic)
	}
	reader.Uint16(&c.Minor).
		Uint16(&c.Major).
		Uint16(&c.ConstantPoolCount)

	return reader.Finish()
}
