package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

// GetHash gets the specified hash of the specified file and returns it.
func GetHash(hashType string, filePath string) string {
	if hashType == "md5" {
		return calcHash(md5.New(), filePath)
	} else if hashType == "sha1" {
		return calcHash(sha1.New(), filePath)
	} else if hashType == "sha256" {

	}
	return ""
}

func calcHash(hash hash.Hash, filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
