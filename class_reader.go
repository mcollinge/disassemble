package disassemble

import (
	"encoding/binary"
	"fmt"

	"github.com/mcollinge/disassemble/utils"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type classinfo struct {
	magic             uint32
	minor             uint16
	major             uint16
	constantPoolCount uint16
	constantPool      []constantinfo
	access            uint16
	name              uint16
	super             uint16
	interfaceCount    uint16
	interfaces        []uint16
	fieldCount        uint16
	fields            []classmemberinfo
	methodCount       uint16
	methods           []classmemberinfo
	attributeCount    uint16
	attributes        []attributeinfo
}

type constantinfo struct {
	tag  uint8
	data []byte
}

type classmemberinfo struct {
	access          uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributeCount  uint16
	attributes      []attributeinfo
}

type attributeinfo struct {
	nameIndex       uint16
	attributeLength uint32
	info            []byte
}

func ReadClass(bytes []byte) error {
	c := &classinfo{}
	reader := utils.NewReader(bytes, binary.BigEndian)

	reader.Uint32(&c.magic)
	if c.magic != 0xCAFEBABE {
		return fmt.Errorf("Magic does not equal 0xCAFEBABE, this files magic is ", c.magic)
	}
	reader.Uint16(&c.minor).
		Uint16(&c.major)

	readConstantPool(reader, c)

	reader.Uint16(&c.access).
		Uint16(&c.name).
		Uint16(&c.super).
		Uint16(&c.interfaceCount)

	c.interfaces = make([]uint16, c.interfaceCount)
	for i := 0; i < int(c.interfaceCount); i++ {
		reader.Uint16(&c.interfaces[i])
	}

	reader.Uint16(&c.fieldCount)
	c.fields = make([]classmemberinfo, c.fieldCount)
	for i := 0; i < int(c.fieldCount); i++ {
		field := readClassMember(reader)
		c.fields[i] = field
	}

	reader.Uint16(&c.methodCount)
	c.methods = make([]classmemberinfo, c.methodCount)
	for i := 0; i < int(c.methodCount); i++ {
		method := readClassMember(reader)
		c.methods[i] = method
	}

	reader.Uint16(&c.attributeCount)
	c.attributes = make([]attributeinfo, c.attributeCount)
	for i := 0; i < int(c.attributeCount); i++ {
		attr := readAttribute(reader)
		c.attributes[i] = attr
	}
	return reader.Error()
}

func readConstantPool(reader *utils.BinaryReader, c *classinfo) error {
	reader.Uint16(&c.constantPoolCount)

	c.constantPool = make([]constantinfo, c.constantPoolCount)
	c.constantPool[0] = constantinfo{}
	for i := 1; i < int(c.constantPoolCount); i++ {
		constantInfo := readConstant(reader)
		c.constantPool[i] = constantInfo
		if constantInfo.tag == CONSTANT_Long || constantInfo.tag == CONSTANT_Double {
			i++
			c.constantPool[i] = constantInfo
		}
	}
	return nil
}

func readConstant(reader *utils.BinaryReader) constantinfo {
	c := &constantinfo{}
	reader.Uint8(&c.tag)
	switch c.tag {
	case CONSTANT_Fieldref, CONSTANT_Methodref, CONSTANT_InterfaceMethodref,
		CONSTANT_Integer, CONSTANT_Float, CONSTANT_NameAndType, CONSTANT_InvokeDynamic:
		c.data = make([]byte, 4)
		reader.Bytes(&c.data)

	case CONSTANT_Long, CONSTANT_Double:
		c.data = make([]byte, 8)
		reader.Bytes(&c.data)

	case CONSTANT_Utf8:
		var length uint16
		reader.Uint16(&length)
		c.data = make([]byte, int(length))
		reader.Bytes(&c.data)

	case CONSTANT_MethodHandle:
		c.data = make([]byte, 3)
		reader.Bytes(&c.data)

	case CONSTANT_Class, CONSTANT_MethodType, CONSTANT_String:
		c.data = make([]byte, 2)
		reader.Bytes(&c.data)
	}
	return *c
}

func readClassMember(reader *utils.BinaryReader) classmemberinfo {
	c := &classmemberinfo{}

	reader.Uint16(&c.access).
		Uint16(&c.nameIndex).
		Uint16(&c.descriptorIndex).
		Uint16(&c.attributeCount)

	c.attributes = make([]attributeinfo, c.attributeCount)
	for i := 0; i < int(c.attributeCount); i++ {
		attr := readAttribute(reader)
		c.attributes[i] = attr
	}
	return *c
}

func readAttribute(reader *utils.BinaryReader) attributeinfo {
	c := &attributeinfo{}
	reader.Uint16(&c.nameIndex).
		Uint32(&c.attributeLength)

	c.info = make([]byte, int(c.attributeLength))
	reader.Bytes(&c.info)
	return *c
}
