package module

import (
	"io/ioutil"
	"os"
	"strings"
)

type Alphabet struct {
	LetterAscii map[string][8]string
}

func getLetter(text []string, s string) [8]string {
	var tableimprimeur [8]string
	for i := 1; i < 9; i++ {
		for _, letter := range s {
			index := int(letter-32) * 9
			tableimprimeur[i-1] = text[index+i]
		}
	}
	return tableimprimeur
}

func GetAlphabet(file []string) *Alphabet {
	var n = make(map[string][8]string)
	result := &Alphabet{LetterAscii: n}
	for i := ' '; i < '~'; i++ {
		result.LetterAscii[string(i)] = getLetter(file, string(i))
	}
	return result
}

func GetSentence(alpha *Alphabet) (result string) {
	args := os.Args[1]
	sentence := make([][8]string, len(args))
	for i := 0; i < len(sentence); i++ {
		sentence[i] = alpha.LetterAscii[string(args[i])]
	}
	for i := 0; i < len(sentence[i]); i++ {
		for j := 0; j < len(sentence); j++ {
			result += sentence[j][i]
		}
		result += string('\n')
	}
	return
}

func GetAlphabetFile() (result []string) {
	file := "standard.txt"
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "standard":
			file = "standard.txt"
		case "shadow":
			file = "shadow.txt"
		case "thinkertoy":
			file = "thinkertoy.txt"
		}
	}
	temp, _ := os.Open("file/" + file)
	temps, _ := ioutil.ReadAll(temp)
	return strings.Split(string(temps), "\r\n")
}
