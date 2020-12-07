package module

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Menu(alphabet *Alphabet) {
	types, param := getParam()
	switch types {
	case "color":
		color := MapColor()
		fmt.Print(color[strings.ToLower(param)])
		PrintSentence(GetSentence(alphabet, true))
		fmt.Print("\u001b[0m")
	case "output":
		if strings.HasSuffix(param, ".txt") {
			f, err := os.Create("output/" + param)
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
	case "reverse":
		if strings.HasSuffix(param, ".txt") {
		}
	default:
		PrintSentence(GetSentence(alphabet, true))
	}
}

func getParam() (types, result string) {
	for _, i := range os.Args[2:] {
		if strings.HasPrefix(i, "--") {
			if strings.Contains(i, "color") {
				result = i[index(i, "=")+1:]
				types = "color"
			}
			if strings.Contains(i, "output") {
				result = i[index(i, "=")+1:]
				types = "output"
			}
			if strings.Contains(i, "reverse") {
				result = i[index(i, "=")+1:]
				types = "reverse"
			}
		}
	}
	return
}

func index(s string, toFind string) int {
	for i := 0; i < len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
