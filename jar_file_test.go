package disassemble

import (
	"testing"
)

func TestJarFile(t *testing.T) {
	jar, err := Open("D:\\Programming\\Java\\java-applet-macro\\JAM.jar")
	if err != nil {
		println(err)
	}
	length := len(jar.Classes)
	println(length)
}
