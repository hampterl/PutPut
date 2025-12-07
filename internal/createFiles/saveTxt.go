package createFiles

import "os"

func SaveTxt(filename, content string) {
	os.WriteFile(filename, []byte(content), 0644)
}
