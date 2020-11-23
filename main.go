package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Alphabet struct {
	LetterAscii map[string][8]string
}

func main() {
	alpha := getAlphabet()
	i := len(os.Args)
	if i >= 1 {
		sentence := make([][8]string, len(os.Args[1]))
		for i := 0; i < len(sentence); i++ {
			sentence[i] = alpha.LetterAscii[string(os.Args[1][i])]
		}
		printSentence(sentence)
	}
}

func getLetter(text []byte, num int) (result [8]string) {
	line := -1
	i := -1
	for j := 0; j < len(text[1:]); j++ {
		if text[j] == '\n' {
			line++
		}
		if line/9 == num && line != -1 {
			if string(text[j]) != "\n" && i <= 7 {
				result[i] += string(text[j])
			}
			if text[j] == '\n' {
				i++
			}
		}
	}
	return
}

func getAlphabet() *Alphabet {
	var n = make(map[string][8]string)
	n[" "] = [8]string{"   ", "   ", "   ", "   ", "   ", "   ", "   ", "   "}
	result := &Alphabet{LetterAscii: n}
	temp, _ := ioutil.ReadFile("standard.txt")
	for i := ' ' - 31; i < '~'-32; i++ {
		result.LetterAscii[string(i+32)] = getLetter(temp, int(i))

	}
	return result
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
			fmt.Print(string(' '))
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
