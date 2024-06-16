package printascii

import (
	"fmt"
	"os"
	"strings"

	"github.com/johneliud/ASCII-Art-Color/colorlibrary"
)

var (
	tabCharText     string
	newLineCharText string
	splitArguments  []string
)

// HandleUnprintableSequences contains a slice of strings that has various flag sequences not supported by the program. A message is displayed whenever an occurrence of the flag is encountered in an input string.
func HandleUnprintableSequences(inputString string) {
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}

	for _, unprintable := range unprintableSequences {
		if strings.Contains(inputString, unprintable) {
			fmt.Println("Input string contains an unprintable sequence")
			return
		}
	}
}

// HandleForeignInput checks whether an input string contains characters out of the range of 32 - 126 within the ASCII manual.
func HandleForeignInput(inputString string) {
	tabCharText = strings.ReplaceAll(inputString, "\\t", "    ")
	newLineCharText = strings.ReplaceAll(tabCharText, "\\n", "\n")
	splitArguments = strings.Split(newLineCharText, "\n")

	for _, splitArg := range splitArguments {
		for _, char := range splitArg {
			if char < 32 || char > 126 {
				fmt.Println("Input string contains a foreign/unprintable character")
				os.Exit(1)
			}
		}
	}
}

// PrintAscii compares letters present in an input string with its index and finds its ASCII character equivalent in the specified banner file.
func PrintAscii(bannerSlice []string, inputString string) {
	var (
		colorFlag     string
		colorToAccess string
		subString     string
		ansiCode      string
		ansiResetCode string = "\033[0m"
	)

	arguments := os.Args[1:]
	if len(arguments) == 2 || len(arguments) == 3 {
		colorFlag = arguments[0]
		colorToAccess = strings.ToLower(colorFlag[8:])
		subString = arguments[1]
	}

	rgbString := strings.ToLower(arguments[0])
	if len(arguments) > 1 && strings.HasPrefix(rgbString, "--color=rgb(") {
		ansiCode = colorlibrary.RgbToAnsi(colorToAccess)
	} else if len(arguments) > 1 && !strings.HasPrefix(rgbString, "--color=rgb(") {
		ansiCode = colorlibrary.ColorLibrary(colorToAccess)
	}

	if inputString == "\\n" {
		fmt.Println()
		return
	} else if inputString == "" {
		return
	} else if inputString == "\t" {
		fmt.Println("    ")
		return
	}
	HandleUnprintableSequences(inputString)
	HandleForeignInput(inputString)

	for _, text := range splitArguments {
		if text == "" {
			fmt.Println()
			continue
		}

		const asciiHeight = 8

		for i := 0; i < asciiHeight; i++ {
			j := 0
			for j < len(text) {
				startingIndex := int(text[j]-32)*9 + 1

				if len(arguments) > 2 && j+len(subString) <= len(text) && text[j:j+len(subString)] == subString {
					for k := 0; k < len(subString); k++ {
						startingIndex := int(text[j+k]-32)*9 + 1
						fmt.Printf(ansiCode + bannerSlice[startingIndex+i] + ansiResetCode)
					}
					j += len(subString)
				} else if len(arguments) == 2 {
					fmt.Printf(ansiCode + bannerSlice[startingIndex+i] + ansiResetCode)
					j++
				} else {
					fmt.Printf(bannerSlice[startingIndex+i])
					j++
				}
			}
			fmt.Println()
		}
	}
}
