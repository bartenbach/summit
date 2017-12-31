package arguments

import "flag"
import "fmt"
import "os"

// Arguments represents the arguments passed to the program.
type Arguments struct {
	SumType     *string
	FilePath    *string
	ExpectedSum *string
}

// ParseArguments parses the command line arguments for the program.
func ParseArguments() Arguments {
	var sType = flag.String("t", "", "Type of checksum to compare")
	var file = flag.String("f", "", "The file to validate")
	var expected = flag.String("e", "", "The expected checksum")
	flag.Parse()
	if *sType != "md5" {
		fmt.Println("Type '", *sType, "' not supported.")
		os.Exit(1)
	}
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		fmt.Println("File '", *file, "' not found.")
		os.Exit(1)
	}
	args := Arguments{SumType: sType, FilePath: file, ExpectedSum: expected}
	return args
}
