package main

import (
	"./module"
	"os"
)

func main() {
	if len(os.Args) >= 1 {
		alpha := module.GetAlphabet(module.GetAlphabetFile())
		module.Menu(alpha)
	}
}
