package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	i := len(os.Args)
	if i >= 1 {
		temp, _ := ioutil.ReadFile("standard.txt")
		sentence := make([][10]string, 12)
		sentence[0] = printLetter(temp, 'H'-32)
		sentence[1] = printLetter(temp, 'e'-32)
		sentence[2] = printLetter(temp, 'l'-32)
		sentence[3] = printLetter(temp, 'l'-32)
		sentence[4] = printLetter(temp, 'o'-32)
		sentence[5] = printLetter(temp, ' '-32)
		sentence[6] = printLetter(temp, 'W'-32)
		sentence[7] = printLetter(temp, 'o'-32)
		sentence[8] = printLetter(temp, 'r'-32)
		sentence[9] = printLetter(temp, 'l'-32)
		sentence[10] = printLetter(temp, 'd'-32)
		sentence[11] = printLetter(temp, '!'-32)
		printSentence(sentence)
	}
}

func printLetter(text []byte, num int) (result [10]string) {
	ligne := 0
	i := 0
	for _, letter := range text {
		if letter == '\n' {
			ligne++
		}
		if ligne/9 == num {
			result[i] += string(letter)
			if letter == '\n' {
				i++
			}
		}
	}
	return
}

func printSentence(sentences [][10]string) {
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
			fmt.Print(string(' '))
		}
		i++
		fmt.Print(string('\n'))
	}
}

func maxLensWord(s [10]string) (result int) {

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
