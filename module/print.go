package module

import (
	"fmt"
)

var (
	colorRed    = "\033[1;31m%s\033[0m"
	colorGreen  = "\033[1;32m%s\033[0m"
	colorYellow = "\033[1;33m%s\033[0m"
	colorBlue   = "\033[1;34m%s\033[0m"
	colorPurple = "\033[1;35m%s\033[0m"
	colorCyan   = "\033[1;36m%s\033[0m"
	colorGray   = "\033[1;37m%s\033[0m"
	color       = "\033[1;38m%s\033[0m"
)

func PrintSentence(sentences [][8]string) {
	i := 0
	for i < len(sentences[i]) {
		j := 0
		for j < len(sentences) {
			k := 0
			for k < len(sentences[j][i]) {
				if sentences[j][i][k] != '\n' && sentences[j][i][k] != '\r' {
					fmt.Printf(color, string(sentences[j][i][k]))
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
