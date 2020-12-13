package module

import (
	"./align"
	"./color"
	"./output"
	"./reverse"
	"./struct"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Menu(alphabet *_struct.Alphabet) {
	param := getParam()
	if v, found := param["color"]; found {
		mapColor := color.MapColor()
		align.Align(mapColor[strings.ToLower(v)], PrintSentence(GetSentence(alphabet, true)), func() (result string) {
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
			output.Output(v, GetSentence(alphabet, false))
		}
	} else if v, found := param["reverse"]; found {
		if strings.HasSuffix(v, ".txt") {
			f, err := os.Open("output/" + v)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			temp, _ := ioutil.ReadAll(f)
			reverse.Reverse(strings.Split(string(temp), "\r\n"), alphabet)
		}
	} else {
		align.Align("\u001b[38m", PrintSentence(GetSentence(alphabet, true)), func() string {
			return "left"
		}, GetLensCMD())
	}
}

func getParam() map[string]string {
	param := make(map[string]string)
	for _, i := range os.Args[1:] {
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
