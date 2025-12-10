package cli

import (
	"PutPut/internal/crypto"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UseSHA256() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("|1|: Hash text using SHA256\n|2|: Hash Files Using\n\n|0|: Return to main menu\n> ")

	scanner.Scan()

	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			sha256String()
			return
		case "2":
			sha256File()
			return
		case "0":
			fmt.Println("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option! Choose from these options:\n|1|: Hash text using SHA256\n" +
				"|2|: Hash Files Using\n\n|0|: Return to main menu\n> ")
			scanner.Scan()
			enDecode = strings.TrimSpace(scanner.Text())
			break
		}
	}
}

func sha256String() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Text to hash: ")
	scanner.Scan()

	fmt.Print("Hashed text: ")
	encodedText := crypto.SHA256String(scanner.Text())
	fmt.Println(encodedText)

	SaveFile(encodedText)
}

func sha256File() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Filepath: ")
	scanner.Scan()
	fileToHash := strings.TrimSpace(scanner.Text())

	hashedFile, err := crypto.SHA256File(fileToHash)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hashed file: " + hashedFile)

	SaveFile(hashedFile)
}
