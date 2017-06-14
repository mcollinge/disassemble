package disassemble

import (
	"archive/zip"
	"bytes"
	"fmt"
	"strings"

	"github.com/mcollinge/disassemble/definitions"
	"github.com/mcollinge/disassemble/packing"
)

type JarFile struct {
	Classes []*definitions.ClassDefinition
}

func Open(path string) (*JarFile, error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	jarFile := &JarFile{}
	jarFile.Classes = make([]*definitions.ClassDefinition, 0)

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
		classDef := definitions.ClassDefinition{}
		err = packing.Unpack(byteBuffer, &classDef)
		if err != nil {
			return nil, err
		}
		jarFile.Classes = append(jarFile.Classes, &classDef)
	}
	return jarFile, nil
}
