package align

func Align(color string, sentence string, align func() string, sizeCmd int) {
	switch align() {
	case "left":
		left(color, sentence)
	case "right":
		right(color, sentence, sizeCmd)
	case "center":
		center(color, sentence, sizeCmd)
	case "justify":
		justify(color, sentence, sizeCmd)
	default:
		left(color, sentence)
	}
}

func getSpace(num int) (result string) {
	for j := 0; j < num; j++ {
		result += " "
	}
	return
}
