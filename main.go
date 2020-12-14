package main

import (
	"./module"
	"os"
)

func main() {
	param := os.Args
	if len(param) >= 2 {
		param = param[1:]
		module.Start(param)
	} else {
		module.Menu()
	}
}
