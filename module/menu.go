package module

import (
	"io/ioutil"
	"os"
)

func menu() {

}

func GetFile() (result []byte) {
	file := ""
	switch os.Args[2] {
	case "standard":
		file = "Ascii.txt"
	case "shadow":
		file = "shadow.txt"
	case "thinkertoy":
		file = "thinkertoy.txt"
	default:
		file = "Ascii.txt"
	}
	result, _ = ioutil.ReadFile("file/" + file)
	return
}
