package choose

import (
	"PutPut/cli"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChooseHashing() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	input := strings.TrimSpace(scanner.Text())

	for {
		switch input {
		case "1":
			cli.UseSHA256()
			return
		case "2":
			fmt.Println("Not yet implemented")
			return
		case "0":
			fmt.Println("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option! Choose from these options:\n|1|: Aes256\n|2|: Comming Soon\n\n|0|: Back\n> ")
			scanner.Scan()
			input = strings.TrimSpace(scanner.Text())
			break
		}
	}
}
