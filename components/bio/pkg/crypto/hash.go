package crypto

import (
	"encoding/hex"
	"io/ioutil"

	"golang.org/x/crypto/blake2b"
)

// HashFile takes a filepath and converts it to bytes
// It then leverages HashBytes to return a string of the
// hashed file contents
func HashFile(file string) (string, error) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return HashBytes(fileBytes), nil
}

// HashBytes takes a byte array and uses blake2b
// to hash the contents. It then returns a hex encoded
// string of the results
func HashBytes(bytes []byte) string {
	hash := blake2b.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}
