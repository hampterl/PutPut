package enDecoding

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Base64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded
}

func Base64Decode(input string) string {
	decoded, _ := base64.StdEncoding.DecodeString(input)

	return string(decoded)
}

func Base64FileEncode(inputPath, savepath string) error {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("failed to read file!")
		return err
	}
	file = []byte(Base64Encode(string(file)))

	return os.WriteFile(savepath, file, 644)
}

func Base64FileDecode(inputPath, savepath string) error {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("failed to read file!")
		return err
	}
	file = []byte(Base64Decode(string(file)))

	return os.WriteFile(savepath, file, 644)
}
