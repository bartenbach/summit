package arguments

import "flag"
import "fmt"
import "os"

// Arguments represents the arguments passed to the program.
type Arguments struct {
	HashType     *string
	FilePath     *string
	ExpectedHash *string
}

// ParseArguments parses the command line arguments for the program.
func ParseArguments() Arguments {
	var hType = flag.String("t", "", "Hash algorithm to use (md5, sha1)")
	var file = flag.String("f", "", "The path to the file to check")
	var expected = flag.String("e", "", "The expected file hash (optional)")
	flag.Parse()
	if *hType != "md5" {
		fmt.Println("Algorithm '", *hType, "' not yet supported.")
		os.Exit(1)
	}
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		fmt.Println("File '", *file, "' not found.")
		os.Exit(1)
	}
	args := Arguments{HashType: hType, FilePath: file, ExpectedHash: expected}
	return args
}
