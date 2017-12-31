package main

import (
	"com/blakebartenbach/checksum/arguments"
	"com/blakebartenbach/checksum/colors"
	"com/blakebartenbach/checksum/summer"
	"fmt"
)

func main() {
	args := arguments.ParseArguments()
	sum := summer.GetSum(*args.SumType, *args.FilePath)
	if sum != *args.ExpectedSum {
		fmt.Println(colors.RED, "Sums do not match!", colors.RESET)

	} else {
		fmt.Println(colors.GREEN, "Sums match!", colors.RESET)
	}
	fmt.Println("    Expected: ", colors.YELLOW, *args.ExpectedSum, colors.RESET)
	fmt.Println("    Found:    ", colors.YELLOW, sum, colors.RESET)
}
