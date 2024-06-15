package colorlibrary

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ValidateValue accepts three parameters and checks if the values are within the range of 0 - 255. If not, the program ceases operation and exits with an os.Exit(1) code.
func ValidateValue(r, g, b int) {
	if r < 0 || g < 0 || b < 0 || r > 255 || g > 255 || b > 255 {
		fmt.Println("Either of values passed is not within the range of 0 - 255")
		os.Exit(1)
	}
}

// RgbToAnsi obtains a string argument passed from the terminal, splits it on certain conditions to obtain three values, converts the values to integer types and returns a formatted ansi string
func RgbToAnsi(colorToAccess string) string {
	var ansiCode string

	rgbString := strings.ToLower(os.Args[1])

	val1 := strings.Split(rgbString, ",")[0]
	val1 = strings.Split(val1, "(")[1]

	val2 := strings.Split(rgbString, ",")[1]

	val3 := strings.Split(rgbString, ")")[0]
	val3 = strings.Split(val3, ",")[2]

	red := strings.TrimSpace(val1)
	green := strings.TrimSpace(val2)
	blue := strings.TrimSpace(val3)

	r, err := strconv.Atoi(red)
	if err != nil {
		fmt.Println("Failed converting value to integer")
		os.Exit(1)
	}
	g, err := strconv.Atoi(green)
	if err != nil {
		fmt.Println("Failed converting value to integer")
		os.Exit(1)
	}
	b, err := strconv.Atoi(blue)
	if err != nil {
		fmt.Println("Failed converting value to integer")
		os.Exit(1)
	}

	ValidateValue(r, g, b)
	ansiCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return ansiCode
}

// ColorLibrary contains accepts a string and return a string. It compares colorToAccess with colors present in the colorsMap and returns the ansi equivalent.
func ColorLibrary(colorToAccess string) string {
	colorsMap := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"blue":   "\033[34m",
		"orange": "\033[38;5;208m",
		"yellow": "\033[33m",
		"black":  "\033[30m",
		"white":  "\033[37m",
		"pink":   "\033[95m",
		"teal":   "\033[36m",
		"purple": "\033[35m",
		"brown":  "\033[33;2m",
		"beige":  "\033[33;2m",
		"indigo": "\033[94m",
		"violet": "\033[94m",
		"maroon": "\033[31;2m",
		"cream":  "\033[97m",
	}

	ansiCode, ok := colorsMap[colorToAccess]
	if !ok {
		fmt.Println("Unsupported color. Refer to colors listed in the README.md file")
		os.Exit(1)
	}
	return ansiCode
}
