package module

import (
	"os"
	"strings"
)

func SetParam() {
	for _, i := range os.Args[1:] {
		if strings.HasPrefix(i, "--") {
			if strings.Contains(i, "color") {
				Getcolor(i)
			}
		}
	}
}

func Getcolor(s string) {
	if strings.Contains(s, "red") {
		color = colorRed
	} else if strings.Contains(s, "green") {
		color = colorGreen
	} else if strings.Contains(s, "yellow") {
		color = colorYellow
	} else if strings.Contains(s, "blue") {
		color = colorBlue
	} else if strings.Contains(s, "purple") {
		color = colorPurple
	} else if strings.Contains(s, "cyan") {
		color = colorCyan
	} else if strings.Contains(s, "gray") {
		color = colorGray
	}
}
