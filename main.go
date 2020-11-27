package main

import (
	"./module"
	"os"
)

func main() {
	module.SetParam()
	alpha := module.GetAlphabet(module.GetFile())
	if len(os.Args) >= 1 {
		sentence := make([][8]string, len(os.Args[1]))
		for i := 0; i < len(sentence); i++ {
			sentence[i] = alpha.LetterAscii[string(os.Args[1][i])]
		}
		module.PrintSentence(sentence)
	}
}
