package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func SHA256String(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}

func SHA256File(inputPath string) (string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	h := sha256.New()

	if _, err := io.Copy(h, file); err != nil {
		return "", nil
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
