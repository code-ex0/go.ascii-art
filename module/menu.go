package module

import (
	"./align"
	"log"
	"os"
	"strings"
)

func Menu(alphabet *Alphabet) {
	param := getParam()
	if v, found := param["color"]; found {
		color := MapColor()
		align.Align(color[strings.ToLower(v)], PrintSentence(GetSentence(alphabet, true)), func() (result string) {
			if v, found := param["align"]; found {
				result = v
			} else {
				result = "left"
			}
			return
		}, GetLensCMD())
	} else if v, found := param["align"]; found {
		align.Align("\u001b[38m", PrintSentence(GetSentence(alphabet, true)), func() string {
			return v
		}, GetLensCMD())
	} else if v, found := param["output"]; found {
		if strings.HasSuffix(v, ".txt") {
			f, err := os.Create("output/" + v)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			for _, letter := range GetSentence(alphabet, false) {
				_, err2 := f.WriteString(letter)
				if err2 != nil {
					log.Fatal(err2)
				}
			}
		}
	} else if _, found := param["reverse"]; found {
	} else {
		align.Align("\u001b[38m", PrintSentence(GetSentence(alphabet, true)), func() string {
			return "left"
		}, GetLensCMD())
	}
}

func getParam() map[string]string {
	param := make(map[string]string)
	for _, i := range os.Args[2:] {
		if strings.HasPrefix(i, "--") {
			if strings.Contains(i, "color") {
				param["color"] = i[index(i, "=")+1:]
			}
			if strings.Contains(i, "output") {
				param["output"] = i[index(i, "=")+1:]
			}
			if strings.Contains(i, "reverse") {
				param["reverse"] = i[index(i, "=")+1:]
			}
			if strings.Contains(i, "align") {
				param["align"] = i[index(i, "=")+1:]
			}
		}
	}
	return param
}

func index(s string, toFind string) int {
	for i := 0; i < len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
