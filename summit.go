package main

import (
	"com.blakebartenbach/summit/arguments"
	"com.blakebartenbach/summit/hasher"
	"com.blakebartenbach/summit/report"
)

func main() {
	args := arguments.ParseArguments()
	fileHash := hasher.GetHash(*args.HashType, *args.FilePath)
	report.CreateReport(fileHash, args)
}
