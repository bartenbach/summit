package main

import (
	"com/blakebartenbach/summit/arguments"
	"com/blakebartenbach/summit/colors"
	"com/blakebartenbach/summit/hasher"
	"fmt"
)

// VERSION is the current version of this program.
const VERSION = "0.0.1-indev"

func main() {
	args := arguments.ParseArguments()
	fileHash := hasher.GetHash(*args.HashType, *args.FilePath)
	if fileHash != *args.ExpectedHash {
		fmt.Println(colors.RED, "Sums do not match!", colors.RESET)

	} else {
		fmt.Println(colors.GREEN, "Sums match!", colors.RESET)
	}
	fmt.Println("    Expected: ", colors.YELLOW, *args.ExpectedHash, colors.RESET)
	fmt.Println("    Found:    ", colors.YELLOW, fileHash, colors.RESET)
}
