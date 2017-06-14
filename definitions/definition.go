package definitions

import "github.com/mcollinge/disassemble/bit"

type Definition interface {
	Unpack(reader *bit.BinaryReader) error
}
