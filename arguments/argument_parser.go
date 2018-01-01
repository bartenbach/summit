package arguments

import "flag"
import "fmt"
import "os"

// VERSION is the current version of this program.
const VERSION = "0.0.1-indev"

// Arguments represents the arguments passed to the program.
type Arguments struct {
	HashType     *string
	FilePath     *string
	ExpectedHash *string
}

// ParseArguments parses the command line arguments for the program.
func ParseArguments() Arguments {
	var hType = flag.String("t", "", "Hash algorithm to use (md5, sha1, sha256, sha512)")
	var file = flag.String("f", "", "The path to the file to check")
	var expected = flag.String("e", "", "The expected file hash (optional)")
	flag.Parse()
	if *hType == "" {
		fmt.Println("summit version ", VERSION)
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *hType != "md5" && *hType != "sha1" && *hType != "sha256" && *hType != "sha512" {
		fmt.Println("Algorithm '", *hType, "' not yet supported.")
		fmt.Println("(see summit -h for help)")
		os.Exit(1)
	}
	if *file == "" {
		fmt.Println("No input file specified.  Please specify a file using -f")
		fmt.Println("(see summit -h for help)")
		os.Exit(1)
	}
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		fmt.Println("File '", *file, "' not found.")
		os.Exit(1)
	}
	args := Arguments{HashType: hType, FilePath: file, ExpectedHash: expected}
	return args
}
