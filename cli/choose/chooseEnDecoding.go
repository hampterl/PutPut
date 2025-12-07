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

	switch input {
	case "1":
		cli.UseBase64()
		break
	case "2":
		fmt.Print("Coming soon!")
		break
	case "0":
		fmt.Print("Exiting...")
		break
	default:
		fmt.Print("Invalid option")
		break
	}
}
