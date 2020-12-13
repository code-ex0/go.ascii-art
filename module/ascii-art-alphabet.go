package module

import (
	"./struct"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getLetter(text []string, s string) (result [8]string) {
	for i := 1; i < 9; i++ {
		for _, letter := range s {
			index := int(letter-32) * 9
			result[i-1] = text[index+i]
		}
	}
	return
}

func GetAlphabet(file []string) (result *_struct.Alphabet) {
	result = &_struct.Alphabet{LetterAscii: make(map[string][8]string)}
	for i := ' '; i < '~'; i++ {
		result.LetterAscii[string(i)] = getLetter(file, string(i))
	}
	return
}

func GetSentence(alpha *_struct.Alphabet, Multyline bool) (result [8]string) {
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
				result[i] += "\a"
				num++
			}
			result[i] += sentence[j][i]
		}
		result[i] += "\n"
	}
	return
}

func GetLensCMD() (ScrnLen int) {
	Out, Err := exec.Command("ScriptBat\\widthcmd.bat").Output()
	if Err == nil {
		OutS := strings.Split(string(Out), "\r\n")
		if len(OutS) >= 3 {
			ScrnLen, _ = strconv.Atoi(OutS[2])
		}
	}
	if ScrnLen < 48 {
		return 48
	}
	return
}

func PrintSentence(sentence [8]string) (result string) {
	for i := 0; i < strings.Count(sentence[0], "\a")+1; i++ {
		for j := 0; j < len(sentence); j++ {
			temp := strings.Split(sentence[j], "\a")
			result += temp[i]
			if !strings.ContainsAny(temp[i], "\n") {
				result += "\n"
			}
		}
	}
	return
}
