package main

import (
	"strings"
)

func cleanInput(test string) []string {
	output := strings.ToLower(test)
	return strings.Fields(output)
}

func main() {
	startRepl()
}