package module

func MapColor() (a map[string]string) {
	a = make(map[string]string)
	a["red"] = "\u001b[31m"
	a["green"] = "\u001b[32m"
	a["yellow"] = "\u001b[33m"
	a["blue"] = "\u001b[34m"
	a["purple"] = "\u001b[35m"
	a["cyan"] = "\u001b[36m"
	a["gray"] = "\u001b[37m"
	a["white"] = "\u001b[38m"
	return
}
