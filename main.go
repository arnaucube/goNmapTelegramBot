package main

import (
	"strings"
)

func checkIp(text string) bool {
	if len(strings.Split(text, ".")) != 4 {
		return false
	}
	if len(strings.Split(text, "")) > 16 {
		return false
	}

	return true
}

func main() {
	readConfig()
	telegramBot()
}
