package choose

import (
	"PutPut/cli"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChooseEnDecoding() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := strings.TrimSpace(scanner.Text())

	for {
		switch input {
		case "1":
			cli.UseBase64()
			return
		case "2":
			cli.UseHex()
			return
		case "0":
			fmt.Print("Press enter to return to main menu...")
			return
			fmt.Print("problem in choose")
		default:
			fmt.Print("Invalid option")
			scanner.Scan()
			input = strings.TrimSpace(scanner.Text())
			break
		}
	}
}
