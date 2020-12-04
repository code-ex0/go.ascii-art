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
		fmt.Print(color[param], GetSentence(alphabet), "\u001b[0m")
	case "output":
		if strings.HasSuffix(param, ".txt") {
			f, err := os.Create("output/" + param)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err2 := f.WriteString(GetSentence(alphabet))
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	default:
		fmt.Print(GetSentence(alphabet))
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
