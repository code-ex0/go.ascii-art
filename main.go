package main

import (
	"./module"
	"os"
)

func main() {
	if len(os.Args) >= 1 {
		module.Menu(module.GetAlphabet(module.GetAlphabetFile()))
	}
}
