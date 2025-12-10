package cli

import (
	"PutPut/internal/enDecoding"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UseHex() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("|1|: Encode text to Hex\n|2|: Encode File with Hex\n\n|3|: Decode Hex to text\n" +
		"|4|: Decode Hex encoded file\n\n|0|: Return to main menu\n> ")

	scanner.Scan()
	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encodeHex()
			return
		case "2":
			encodehexFile()
			return
		case "3":
			decodeHex()
			return
		case "4":
			decodehexFile()
			return
		case "0":
			fmt.Print("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option! Choose from these options:\n|1|: Encode text to Hex\n|2|: Encode File with Hex\n\n|3|: Decode Hex to text\n" +
				"|4|: Decode Hex encoded file\n\n|0|: Return to main menu\n> ")
			scanner.Scan()
			enDecode = strings.TrimSpace(scanner.Text())
			break
		}
	}
}

func encodeHex() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Text to encode: ")
	scanner.Scan()

	fmt.Print("Encoded text: ")
	encodedText := enDecoding.HexEncode(scanner.Text())
	fmt.Println(encodedText)

	SaveFile(encodedText)
}

func decodeHex() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Hex text: ")
	scanner.Scan()

	fmt.Print("Decoded text: ")
	decodedText := enDecoding.HexDecode(scanner.Text())
	fmt.Println(decodedText)

	SaveFile(decodedText)
}

func encodehexFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Filepath: ")
	scanner.Scan()
	inputpath := strings.TrimSpace(scanner.Text())

	fmt.Print("Save location: ")
	scanner.Scan()
	outputPath := strings.TrimSpace(scanner.Text()) + ".enc"

	enDecoding.HexFileEncode(inputpath, outputPath)

	fmt.Println("File encoded! Press enter to return to main menu...")
}

func decodehexFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Filepath: ")
	scanner.Scan()
	inputpath := strings.TrimSpace(scanner.Text())

	fmt.Print("Save location: ")
	scanner.Scan()
	outputPath := strings.TrimSpace(scanner.Text())

	enDecoding.HexFileDecode(inputpath, outputPath)

	fmt.Println("File decoded! Press enter to return to main menu...")
}
