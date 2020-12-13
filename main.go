package main

import (
	"./module"
	"./module/fs"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		module.Menu(module.GetAlphabet(fs.GetAlphabetFile()))
	}
}
