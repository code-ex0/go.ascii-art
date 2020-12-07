package module

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Alphabet struct {
	LetterAscii map[string][8]string
}

func getLetter(text []string, s string) (result [8]string) {
	for i := 1; i < 9; i++ {
		for _, letter := range s {
			index := int(letter-32) * 9
			result[i-1] = text[index+i]
		}
	}
	return
}

func GetAlphabet(file []string) *Alphabet {
	result := &Alphabet{LetterAscii: make(map[string][8]string)}
	for i := ' '; i < '~'; i++ {
		result.LetterAscii[string(i)] = getLetter(file, string(i))
	}
	return result
}

func GetSentence(alpha *Alphabet, Multyline bool) (result [8]string) {
	args := os.Args[1]
	LensCMD := GetLensCMD()
	sentence := make([][8]string, len(args))
	for i := 0; i < len(sentence); i++ {
		sentence[i] = alpha.LetterAscii[string(args[i])]
	}
	for i := 0; i < len(sentence[i]); i++ {
		num := 1
		for j := 0; j < len(sentence); j++ {
			if (len(sentence[j][i])+len(result[i]))+(num*5) >= num*LensCMD-1 && Multyline {
				result[i] += "\r"
				num++
			}
			result[i] += sentence[j][i]
		}
		result[i] += "\n"
	}
	return
}

func GetLensCMD() int {
	Out, Err := exec.Command("ScriptBat\\widthcmd.bat").Output()
	var ScrnLen int
	if Err == nil {
		OutS := strings.Split(string(Out), "\r\n")
		if len(OutS) >= 3 {
			ScrnLen, _ = strconv.Atoi(OutS[2])
		}
	}
	return ScrnLen
}

func PrintSentence(sentence [8]string) {
	for i := 0; i < strings.Count(sentence[0], "\r")+1; i++ {
		for j := 0; j < len(sentence); j++ {
			temp := strings.Split(sentence[j], "\r")
			fmt.Print(temp[i])
			if !strings.ContainsAny(temp[i], "\n") {
				fmt.Println()
			}
		}
	}
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
