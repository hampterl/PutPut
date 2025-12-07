package cli

import (
	"PutPut/internal/createFiles"
	"PutPut/internal/enDecoding"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UseBase64() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		fmt.Println("Scanner error!")
	}

	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encodebase()
			return
		case "2":
			decodebase()
			return
		default:
			fmt.Println("Invalid option\n|1|: Encode text to base64\n|2|: Decode base64 to text\n>")
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
		}
	}
}

func decodebase() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Base64 text: ")
	scanner.Scan()

	fmt.Print("Decoded text: ")
	decodedText := enDecoding.Base64Decode(scanner.Text())
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
