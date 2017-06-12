package definitions

import "github.com/mcollinge/disassemble/bit"

type Definition interface {
	Unpack(unpacker *bit.BinaryReader)
}
