package main

import (
	"./module"
	"fmt"
	"os"
)

func main() {

	alpha := module.GetAlphabet(module.GetFile())
	i := len(os.Args)
	if i >= 1 {
		sentence := make([][8]string, len(os.Args[1]))
		for i := 0; i < len(sentence); i++ {
			sentence[i] = alpha.LetterAscii[string(os.Args[1][i])]
		}
		printSentence(sentence)
	}
}

func printSentence(sentences [][8]string) {
	i := 0
	for i < len(sentences[i]) {
		j := 0
		for j < len(sentences) {
			k := 0
			for k < len(sentences[j][i]) {
				if sentences[j][i][k] != '\n' && sentences[j][i][k] != '\r' {
					fmt.Print(string(sentences[j][i][k]))
				} else if sentences[j][i][k] == '\r' {
					for space := 0; space < maxLensWord(sentences[j])-k; space++ {
						fmt.Print(" ")
					}
				}
				k++
			}
			j++
		}
		i++
		fmt.Print(string('\n'))
	}
}

func maxLensWord(s [8]string) (result int) {
	for _, word := range s {
		maxLens := 0
		for _, letter := range word {
			if letter >= ' ' && letter <= '~' {
				maxLens++
			}
		}
		if maxLens > result {
			result = maxLens
		}
	}
	return
}
