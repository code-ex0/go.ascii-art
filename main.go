package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Alphabet struct {
	LetterAscii [7]string
	Letter      string
	Next        *Alphabet
}

func main() {

	alpha := getLetter()
	i := len(os.Args)
	if i >= 1 {
		sentence := make([][7]string, len(os.Args[1]))
		for i := 0; i < len(sentence); i++ {
			it := alpha
			for it != nil {
				if string(os.Args[1][i]) == it.Letter {
					sentence[i] = it.LetterAscii
				}
				it = it.Next
			}
		}
		printSentence(sentence)
	}
}

func printLetter(text []byte, num int) (result [7]string) {
	line := -1
	i := -1
	for j := 0; j < len(text[1:]); j++ {
		if text[j] == '\n' {
			line++
		}
		if line/9 == num && line != -1 {
			if string(text[j]) != "\n" && i <= 6 {
				result[i] += string(text[j])
			}
			if text[j] == '\n' {
				i++
			}
		}
	}
	return
}

func getLetter() *Alphabet {
	space := [7]string{"   ", "   ", "   ", "   ", "   ", "   ", "   "}
	result := &Alphabet{LetterAscii: space, Letter: " "}
	temp, _ := ioutil.ReadFile("standard.txt")
	it := result
	for i := ' ' - 31; i < '~'-32; i++ {
		n := &Alphabet{LetterAscii: printLetter(temp, int(i)), Letter: string(i + 32)}
		it.Next = n
		it = it.Next
	}
	return result
}

func printSentence(sentences [][7]string) {
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

func maxLensWord(s [7]string) (result int) {

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
