package main

import "fmt"

func main() {
	fmt.Print(numberSpace(2, 186, 124))
}

func numberSpace(numbersSpace int, sizeCmd int, sizeSentence int) []int {
	var result = make([]int, numbersSpace)
	allSpace := sizeCmd - sizeSentence
	for i := 0; i < numbersSpace; i++ {
		temp := allSpace / numbersSpace
		if allSpace%numbersSpace != 0 {
			result[i] = temp + 1
			allSpace--
		} else {
			result[i] = temp
		}
	}
	return result
}
