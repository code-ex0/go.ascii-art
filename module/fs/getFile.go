package fs

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetAlphabetFile() (result []string) {
	file := "standard.txt"
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "standard":
			file = "standard.txt"
		case "shadow":
			file = "shadow.txt"
		case "thinkertoy":
			file = "thinkertoy.txt"
		}
	}
	temp, _ := os.Open("file/" + file)
	temps, _ := ioutil.ReadAll(temp)
	return strings.Split(string(temps), "\r\n")
}
