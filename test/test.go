package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	param := os.Args
	if len(param) == 3 {
		x, _ := strconv.Atoi(param[1])
		y, _ := strconv.Atoi(param[2])
		if !(y <= 0 || x <= 0) {
			getHead(x)
			getBody(x, y)
			if y >= 2 {
				getHead(x)
			}
		} else {
			fmt.Println("not possible...")
		}
	} else {
		fmt.Println("incomplete...")
	}
}

func getHead(sizex int) {
	fmt.Print("x")
	if sizex > 1 {
		for i := 0; i < sizex-2; i++ {
			fmt.Print("-")
		}
		fmt.Print("x")
	}
	fmt.Print("\n")
}

func getBody(sizex, sizey int) {
	k := sizey - 2
	for i := 0; i < sizey-2; i++ {
		fmt.Print(i)
		if sizex > 1 {
			for j := 0; j < sizex-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print(k)
			k--
		}
		fmt.Print("\n")
	}
}
