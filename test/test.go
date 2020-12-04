package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("ERROR: Too many arguments!")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("ERROR: Not enough arguments!")
		return
	}

	str := os.Args[1]

	file, err := os.Open("ascii/standard.txt")
	ErrCheck(err)

	data, err := ioutil.ReadAll(file)
	ErrCheck(err)

	for _, letter := range str {
		if letter < 32 || letter > 126 {
			fmt.Println("ERROR: Non ascii character detected")
			os.Exit(1)
		}
	}

	res := strings.Split(string(data), "\r\n")

	var tableaufinal [][8]string

	for _, v := range str {
		tableaufinal = append(tableaufinal, PrintWord(string(v), res))
	}

	for i := 0; i < len(tableaufinal[i]); i++ {
		for g := 0; g < len(tableaufinal); g++ {
			fmt.Print(tableaufinal[g][i])
		}
		fmt.Println()
	}
}

func PrintWord(word string, tab []string) [8]string {
	var tableimprimeur [8]string
	for i := 1; i < 9; i++ {
		for _, letter := range word {
			index := int(letter-32) * 9
			tableimprimeur[i-1] = tab[index+i]
		}
	}
	return tableimprimeur
}

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
