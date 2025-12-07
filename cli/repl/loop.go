package repl

import (
	"PutPut/cli/choose"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Loop() {
	scanner := bufio.NewScanner(os.Stdin)

	//loop until exit
	for {
		//Start Menu Display
		Menu()

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("|1|: Base64\n|2|: Hex\n|0|: Back\n> ")
			choose.ChooseEnDecoding()
		case "2":
			fmt.Print("Comming soon!")
		case "0":
			fmt.Print("Exiting...")
			return
		default:
			fmt.Print("Invalid option")
		}
		scanner.Scan()
	}
}
