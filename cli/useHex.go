package cli

import (
	"PutPut/internal/createFiles"
	"PutPut/internal/enDecoding"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UseHex() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("|1|: Encode text to Hex\n|2|: Decode Hex to text\n|0|: Return to main menu\n> ")

	scanner.Scan()
	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encodeHex()
			return
		case "2":
			decodeHex()
			return
		case "0":
			fmt.Print("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option\n|1|: Encode text to Hex\n|2|: Decode Hex to text\n> ")
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

	fmt.Print("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())
	for {
		switch saveFile {
		case "y":
			fmt.Print("Filename: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())

			createFiles.SaveTxt(filename, encodedText)
			fmt.Println("File saved!\nPress enter to return to main menu...")
			return

		case "n":
			fmt.Println("NO file saved!\nPress enter to return to main menu...")
			return

		default:
			fmt.Print("Please enter y or n\nSave txt into file? (y/n): ")
			scanner.Scan()
			saveFile = strings.TrimSpace(scanner.Text())
		}
	}
}

func decodeHex() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Base64 text: ")
	scanner.Scan()

	fmt.Print("Decoded text: ")
	decodedText := enDecoding.HexDecode(scanner.Text())
	fmt.Println(decodedText)

	fmt.Print("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())
	for {
		switch saveFile {
		case "y":
			fmt.Print("Filename: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())

			createFiles.SaveTxt(filename, decodedText)
			fmt.Println("File saved!\nPress enter to return to main menu...")
			return

		case "n":
			fmt.Println("NO file saved!\nPress enter to return to main menu...")
			return

		default:
			fmt.Print("Please enter y or n\nSave txt into file? (y/n): ")
			scanner.Scan()
		}
	}
}
