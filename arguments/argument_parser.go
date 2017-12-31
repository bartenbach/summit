package arguments

import "flag"

// Arguments represents the arguments passed to the program.
type Arguments struct {
	SumType     *string
	FilePath    *string
	ExpectedSum *string
}

// ParseArguments parses the command line arguments for the program.
func ParseArguments() Arguments {
	var sum = flag.String("t", "", "Type of checksum to compare")
	var file = flag.String("f", "", "The file to validate")
	var expected = flag.String("e", "", "The expected checksum")
	flag.Parse()
	args := Arguments{SumType: sum, FilePath: file, ExpectedSum: expected}
	return args
}
