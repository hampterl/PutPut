package cli

import (
	"PutPut/internal/enDecoding"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UseBase64() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("|1|: Encode text with Base64\n|2|: Encode File with Base64\n\n|3|: Decode Base64 to text" +
		"\n|4|: Decode Base64 encoded file\n\n|0|: Return to main menu\n> ")

	scanner.Scan()
	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encodebase()
			return
		case "2":
			encodebaseFile()
			return
		case "3":
			decodebase()
			return
		case "4":
			decodebaseFile()
			return
		case "0":
			fmt.Print("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option! Choose from these options:\n|1|: Encode text to base64\n" +
				"|2|: Encode File with Base64\n\n|3|: Decode Base64 to text\n|4|: Decode Base64 encoded file\n\n" +
				"|0|: Return to main menu\n> ")
			scanner.Scan()
			enDecode = strings.TrimSpace(scanner.Text())
			break
		}
	}
}

func encodebase() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Text to encode: ")
	scanner.Scan()

	fmt.Print("Encoded text: ")
	encodedText := enDecoding.Base64Encode(scanner.Text())
	fmt.Println(encodedText)

	SaveFile(encodedText)
}

func decodebase() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Base64 text: ")
	scanner.Scan()

	fmt.Print("Decoded text: ")
	decodedText := enDecoding.Base64Decode(scanner.Text())
	fmt.Println(decodedText)

	SaveFile(decodedText)
}

func encodebaseFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Filepath: ")
	scanner.Scan()
	inputpath := strings.TrimSpace(scanner.Text())

	fmt.Print("Save location: ")
	scanner.Scan()
	outputPath := strings.TrimSpace(scanner.Text()) + ".enc"

	enDecoding.Base64FileEncode(inputpath, outputPath)

	fmt.Println("File encoded! Press enter to return to main menu...")
}

func decodebaseFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Filepath: ")
	scanner.Scan()
	inputpath := strings.TrimSpace(scanner.Text())

	fmt.Print("Save location: ")
	scanner.Scan()
	outputPath := strings.TrimSpace(scanner.Text())

	enDecoding.Base64FileDecode(inputpath, outputPath)

	fmt.Println("File decoded! Press enter to return to main menu...")
}
