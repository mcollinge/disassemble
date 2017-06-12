package disassemble

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/mcollinge/disassemble/bit"
	"github.com/mcollinge/disassemble/definitions"
)

type JarFile struct {
	//Classes []*ClassDefinition
}

func Open(path string) (*JarFile, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	//jarFile := &JarFile{}
	//jarFile.Classes = []*ClassDefinition{}

jarReader:
	for _, f := range reader.File {
		if !strings.HasSuffix(f.Name, ".class") {
			continue jarReader
		}
		zippedReader, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("Failed to open file in jar. Internal error: %v", err)
		}
		defer zippedReader.Close()

		byteBuffer := new(bytes.Buffer)
		read, err := byteBuffer.ReadFrom(zippedReader)
		if err != nil {
			return nil, err
		}
		if read == 0 {
			return nil, fmt.Errorf("Could not read file: %s", f.Name)
		}
		unpacker := bit.NewReader(byteBuffer, binary.BigEndian)
		classDef := definitions.ClassDefinition{}
		err = classDef.Unpack(unpacker)
		println(f.Name)
		println(classDef.Magic)
		println(classDef.Major)
		println(classDef.ConstantPoolCount)
		println()

	}
	return nil, nil
}
