package cli

import (
	"PutPut/internal/createFiles"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SaveFile(input string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Save txt into file? (y/n): ")
		scanner.Scan()
		saveFile := strings.TrimSpace(scanner.Text())
		for {
			switch saveFile {
			case "y":
				fmt.Print("Filename: ")
				scanner.Scan()
				filename := strings.TrimSpace(scanner.Text())

				createFiles.SaveTxt(filename, input)
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
}
