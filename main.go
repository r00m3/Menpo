package main

import (
	"fmt"
	authenticated "menpo/authenticated_encryption"
	"menpo/colors"
	"menpo/hash"
	"os"
)

var fileName string
var fileAction string

func clearTerminal() {
	fmt.Printf("\x1b[H")
	fmt.Printf("\x1b[2J")
	fmt.Printf("\x1b[3J")
}

func selectFile() {
	clearTerminal()
	colors.GreenBold()
	fmt.Printf("\n    Make sure file and program are in same directory.\n")
	colors.Reset()
	fmt.Printf("    Select file:\n")
	// Scan for file name.
	fmt.Scan(&fileName)
	// Check if file exist.
	file, err := os.Open(fileName)
	if err != nil {
		selectFile()
	}
	defer file.Close()
}

func selectAction() {
	clearTerminal()
	fmt.Printf("\n    Select action for \x1b[36m%v\x1b[0m:\n", fileName)
	fmt.Printf("\x1b[36m[ 1 ]\x1b[0m Calculate sha256sum.\n")
	fmt.Printf("\x1b[36m[ 2 ]\x1b[0m Encrypt with authenticated encryption.\n")
	fmt.Printf("\x1b[36m[ 3 ]\x1b[0m Decrypt with authenticated encryption.\n")
	fmt.Scan(&fileAction)
	switch fileAction {
	case "1":
		hash.CalcSha256(fileName)
	case "2":
		authenticated.AuthEncrypt(fileName)
	case "3":
		authenticated.AuthDecrypt(fileName)
	default:
		selectAction()
	}
}

func main() {
	selectFile()
	selectAction()
}
