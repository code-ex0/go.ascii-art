package module

func MapColor() (a map[string]string) {
	a = make(map[string]string)
	a["red"] = "\033[1;31m%s\033[0m"
	a["green"] = "\033[1;32m%s\033[0m"
	a["yellow"] = "\033[1;33m%s\033[0m"
	a["blue"] = "\033[1;34m%s\033[0m"
	a["purple"] = "\033[1;35m%s\033[0m"
	a["cyan"] = "\033[1;36m%s\033[0m"
	a["gray"] = "\033[1;37m%s\033[0m"
	return
}
