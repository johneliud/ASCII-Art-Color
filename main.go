package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/johneliud/ASCII-Art-Color/printascii"
	"github.com/johneliud/ASCII-Art-Color/readfile"
)

func main() {
	var inputString string

	arguments := os.Args[1:]
	if len(arguments) > 1 && !strings.HasPrefix(arguments[0], "--color=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	switch len(arguments) {
	case 1:
		inputString = arguments[0]
	case 2:
		inputString = arguments[1]
	case 3:
		inputString = arguments[2]
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}

	banner := readfile.ReadFile("standard.txt")
	printascii.PrintAscii(banner, inputString)
}
