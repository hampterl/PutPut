package enDecoding

import (
	"encoding/hex"
	"fmt"
	"os"
)

func HexEncode(input string) string {
	encoded := hex.EncodeToString([]byte(input))

	return encoded
}

func HexDecode(input string) string {
	decoded, _ := hex.DecodeString(input)

	return string(decoded)
}

func HexFileEncode(inputPath, savepath string) error {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("failed to read file!")
		return err
	}
	file = []byte(HexEncode(string(file)))

	return os.WriteFile(savepath, file, 644)
}

func HexFileDecode(inputPath, savepath string) error {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("failed to read file!")
		return err
	}
	file = []byte(HexDecode(string(file)))

	return os.WriteFile(savepath, file, 644)
}
