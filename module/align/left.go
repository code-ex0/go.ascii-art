package align

import "fmt"

func left(color string, sentence string) {
	fmt.Print(color, sentence, "\u001b[0m")
}
