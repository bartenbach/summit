package summer

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

// GetSum gets the specified sum of the specified file and returns it.
func GetSum(sumType string, filePath string) string {
	if sumType == "md5" {
		return calcMd5Sum(filePath)
	} else if sumType == "sha1" {

	} else if sumType == "sha256" {

	}
	return ""
}

func calcMd5Sum(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
