package cli

import (
	"PutPut/internal/crypto"
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func UseAes256() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("|1|: Encrypt text with aes256\n|2|: Encrypt File with aes256\n\n|3|: Decrypt aes256 to text\n" +
		"|4|: Decrypt File with aes256\n\n|0|: Return to main menu\n> ")

	scanner.Scan()

	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encryptAes()
			return
		case "2":
			encryptFile()
			return
		case "3":
			decryptAes()
			return
		case "4":
			decryptFile()
			return
		case "0":
			fmt.Println("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option! Choose from these options:\n|1|: Encrypt text with aes256\n|2|: |2|: Encrypt File with aes256" +
				"|3|: Decrypt aes256 to text\n|4|: Decrypt File with aes256\n\n|0|: Return to main menu\n> ")
			scanner.Scan()
			enDecode = strings.TrimSpace(scanner.Text())
			break
		}
	}
}

func encryptAes() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Print("Text to encrypt: ")
	scanner.Scan()
	plaintext := strings.TrimSpace(scanner.Text())

	fmt.Print("Encrypted text (hex): \n> ")
	ciphertext, err := crypto.EncryptAes(plaintext, []byte(key))
	if err != nil {
		fmt.Println(err)
	}

	hexEncodedAes := hex.EncodeToString(ciphertext)

	fmt.Println(hexEncodedAes)

	SaveFile(hexEncodedAes)
}

func decryptAes() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Print("Encrypted text (hex): ")
	scanner.Scan()
	hexciphertext := strings.TrimSpace(scanner.Text())

	ciphertext, _ := hex.DecodeString(hexciphertext)

	fmt.Print("Decrypted text: ")
	plaintext, err := crypto.DecryptAes(ciphertext, []byte(key))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(plaintext)

	SaveFile(plaintext)
}

func encryptFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n-----------------------------------------------------------------------\n")
	fmt.Print("Instructions in README.md!\n")
	fmt.Print("DISCLAIMER AND WARNING:\nThis WILL encrypt the file you chose! IF you lose your key and overwrite your original, " +
		"you will NOT be able to decrypt the file!\n" +
		"Never share your key with anyone! Make a backup before encrypting! Use at own Risk!\n")
	fmt.Println("\n-----------------------------------------------------------------------\n\n")

	fmt.Print("Filepath: ")
	scanner.Scan()
	filepath := strings.TrimSpace(scanner.Text())

	fmt.Print("Save location: ")
	scanner.Scan()
	savepath := strings.TrimSpace(scanner.Text()) + ".enc"

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	crypto.EncryptFileAes(filepath, savepath, []byte(key))

	fmt.Println("File encrypted! SAVE YOUR KEY!\nPress enter to return to main menu...")
}

func decryptFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n-----------------------------------------------------------------------\n")
	fmt.Print("Instructions in README.md!\n")
	fmt.Print("You are about to decrypt a file! Make sure you use the right key or it won't work correctly!\n")
	fmt.Println("\n-----------------------------------------------------------------------\n\n")

	fmt.Print("Filepath: ")
	scanner.Scan()
	filepath := strings.TrimSpace(scanner.Text()) + ".enc"

	fmt.Print("Outputpath: ")
	scanner.Scan()
	outputPath := strings.TrimSpace(scanner.Text())

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	crypto.DecryptFileAes(filepath, outputPath, []byte(key))

	fmt.Println("File decrypted! SAVE YOUR KEY!\nPress enter to return to main menu...")
}
