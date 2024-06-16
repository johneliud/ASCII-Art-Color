package readfile

import (
	"fmt"
	"os"
	"strings"
)
// ReadFile accepts a string and returns a slice of string. It reads a file, checks for potential file corruptions, splits it depending on the read file and returns the split file.
func ReadFile(banner string) []string {
	bannerFile, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading banner file", err)
		os.Exit(1)
	}

	var splitBannerFile []string

	if banner == "thinkertoy.txt" {
		splitBannerFile = strings.Split(string(bannerFile), "\r\n")
	} else {
		splitBannerFile = strings.Split(string(bannerFile), "\n")
	}

	// os.Stat() contains various properties of a file.
	bannerFileDetails, err := os.Stat(banner)
	if err != nil {
		fmt.Println("Unable to obtain file details of banner file", err)
		os.Exit(1)
	}
	// bannerFileSize stores the size of each banner file and compares it against the expected size to determine if a file has been tampered with/modified.
	bannerFileSize := bannerFileDetails.Size()
	if bannerFileSize != 6623 && bannerFileSize != 7463 && bannerFileSize != 5558 {
		fmt.Println("Corrupted banner file")
		os.Exit(1)
	}
	return splitBannerFile
}
