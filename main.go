package main

import (
	"./module"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		module.Menu(module.GetAlphabet(module.GetAlphabetFile()))
	}
}
