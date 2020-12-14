package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	i := len(os.Args)
	if i >= 1 {
		for _, file := range os.Args[1:] {
			temp, err := ioutil.ReadFile("output/" + file)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
			}
			fmt.Println(string(temp))
		}
	}
}
