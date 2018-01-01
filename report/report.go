package report

import (
	"github.com/proxa/summit/arguments"
	"github.com/proxa/summit/colors"
	"fmt"
	"os"
)

// CreateReport creates a command line report of completed operations.
func CreateReport(fileHash string, args arguments.Arguments) {
	// if no hash was expected, just print the file hash
	if *args.ExpectedHash == "" {
		fmt.Println(colors.CYAN, *args.HashType, colors.RESET)
		fmt.Println("    ", *args.FilePath)
		fmt.Println("    ", fileHash)
		os.Exit(0)
	} else {
		if fileHash != *args.ExpectedHash {
			fmt.Println(colors.RED, "Sums do not match!", colors.RESET)
			printHashDiff(*args.ExpectedHash, fileHash, colors.YELLOW)
			os.Exit(1)
		} else {
			fmt.Println(colors.GREEN, "Sums match!", colors.RESET)
			printHashDiff(*args.ExpectedHash, fileHash, colors.GREEN)
			os.Exit(0)
		}
	}
}

func printHashDiff(expected string, actual string, color string) {
	fmt.Println("    Expected: ", color, expected, colors.RESET)
	fmt.Println("    Found:    ", color, actual, colors.RESET)
}
