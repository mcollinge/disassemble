package disassemble

import (
	"io/ioutil"
	"testing"

	"os"
)

func TestJarFile(t *testing.T) {
	name := "D:\\Programming\\Java\\07kit\\target\\classes\\com\\kit\\Application.class"
	file, err := os.Open(name)
	if err != nil {
		t.Error("Could not load file: ", name)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("Could not read bytes from file: ", name)
	}
	_, err = ReadClass(bytes)
}
