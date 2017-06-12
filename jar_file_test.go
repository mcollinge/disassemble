package disassemble

import (
	"testing"
)

func TestJarFile(t *testing.T) {
	_, err := Open("D:\\Programming\\Java\\java-applet-macro\\JAM.jar")
	if err != nil {
		println(err)
	}
	//println(len(jarFile.Classes))
}
