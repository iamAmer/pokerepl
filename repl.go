package main

import (
	"fmt"
	"bufio"
	"os"
)
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Printf(">> ")

		scanner.Scan()
		text := scanner.Text()

		fmt.Printf("Echoing: %q\n", text)
	}
}