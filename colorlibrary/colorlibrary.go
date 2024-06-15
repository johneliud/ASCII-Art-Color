package colorlibrary

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateValue(r, g, b int) {
	if r < 0 || g < 0 || b < 0 || r > 255 || g > 255 || b > 255 {
		fmt.Println("Either of values passed is not within the range of 0 - 255")
		os.Exit(1)
	}
}

func RgbToAnsi() string {
	var ansiCode string

	arguments := os.Args[1:]
	values := arguments[0]

	val1 := strings.Split(values, ",")[0]
	val1 = strings.Split(val1, "(")[1]

	val2 := strings.Split(values, ",")[1]

	val3 := strings.Split(values, ")")[0]
	val3 = strings.Split(val3, ",")[2]

	red, green, blue := strings.TrimSpace(val1), strings.TrimSpace(val2), strings.TrimSpace(val3)

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

	fmt.Printf("Red: %v - %T\n", r, r)
	fmt.Printf("Green: %v - %T\n", g, g)
	fmt.Printf("Blue: %v - %T\n", b, b)

	return ansiCode
}
