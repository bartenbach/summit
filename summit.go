package main

import (
	"github.com/proxa/summit/arguments"
	"github.com/proxa/summit/hasher"
	"github.com/proxa/summit/report"
)

func main() {
	args := arguments.ParseArguments()
	fileHash := hasher.GetHash(*args.HashType, *args.FilePath)
	report.CreateReport(fileHash, args)
}
