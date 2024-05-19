package authenticated

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"fmt"
	"io"
	"menpo/colors"
	"os"
	"strings"
)

// encrypted_file.txt -> file.txt
func removeUnderscore(name string) string {
	_, removed, _ := strings.Cut(name, "_")
	return removed
}

func AuthEncrypt(name string) {
	// Open file.txt
	file, err := os.Open(name)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error opening %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	defer file.Close()
	// Generate crypto-random 16 byte key.
	key := make([]byte, 16)
	if _, err := io.ReadFull(crand.Reader, key); err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error generating secure key.\n")
		colors.Reset()
		fmt.Println(err)
	}
	// Read from file.
	plaintext, err := io.ReadAll(file)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error reading from %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	// AES block.
	block, err := aes.NewCipher(key)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error constructing AES block.\n")
		colors.Reset()
		fmt.Println(err)
	}
	// Create crypto-random 12 byte nonce.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(crand.Reader, nonce); err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error creating nonce.\n")
		colors.Reset()
		fmt.Println(err)
	}
	// Construct cipher.
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error constructing cipher.\n")
		colors.Reset()
		fmt.Println(err)
	}
	// Encrypt text.
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	// Write encrypted data to encrypted_file.txt
	err = os.WriteFile("encrypted_"+name, ciphertext, 0660)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error writing ciphertext to encrypted_%v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	// Write key to key_file.txt
	err = os.WriteFile("key_"+name, key, 0660)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error writing key to key_%v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	// Write nonce to nonce_file.txt
	err = os.WriteFile("nonce_"+name, nonce, 0660)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error writing nonce to nonce_%v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	fmt.Printf("\nEncrypted data written to \x1b[36mencrypted_%v\x1b[0m", name)
	fmt.Printf("\nSecret key written to \x1b[36mkey_%v\x1b[0m", name)
	fmt.Printf("\nNonce written to \x1b[36mnonce_%v\x1b[0m", name)
	colors.GreenBold()
	fmt.Printf("\n    Do NOT modify or rename those files.\n")
	fmt.Printf("\nExiting.\n")
	colors.Reset()
}

func AuthDecrypt(name string) {
	// encrypted_file.txt -> file.txt
	clearedName := removeUnderscore(name)
	// Open files to read from.
	// Open encrypted_file.txt
	ciphertextFile, err := os.Open(name)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error opening %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	defer ciphertextFile.Close()
	// Open key_file.txt
	keyFile, err := os.Open("key_" + clearedName)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error opening key_%v.\n", clearedName)
		colors.Reset()
		fmt.Println(err)
	}
	defer keyFile.Close()
	// Open nonce_file.txt
	nonceFile, err := os.Open("nonce_" + clearedName)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error opening nonce_%v.\n", clearedName)
		colors.Reset()
		fmt.Println(err)
	}
	defer nonceFile.Close()
	// Read []byte from opened files.
	// CiphertextFile -> []byte ciphertext.
	ciphertext, err := io.ReadAll(ciphertextFile)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error reading from %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	// KeyFile -> []byte key.
	key, err := io.ReadAll(keyFile)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error reading from key_%v.\n", clearedName)
		colors.Reset()
		fmt.Println(err)
	}
	// NonceFile -> []byte nonce.
	nonce, err := io.ReadAll(nonceFile)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error reading from nonce_%v.\n", clearedName)
		colors.Reset()
		fmt.Println(err)
	}
	// Construct AES cipher block.
	block, err := aes.NewCipher(key)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error constructing AES block.\n")
		colors.Reset()
		fmt.Println(err)
	}
	// Construct AES cipher.
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error constructing AES cipher.\n")
		colors.Reset()
		fmt.Println(err)
	}
	// Try opening encrypted file using data from opened files.
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error decrypting %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	// Write plaintext to decrypted file.
	err = os.WriteFile("decrypted_"+clearedName, plaintext, 0660)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error writing decrypted data to decrypted_%v.\n", clearedName)
		colors.Reset()
		fmt.Println(err)
	}
	fmt.Printf("\nDecrypted data written to \x1b[36mdecrypted_%v\x1b[0m", clearedName)
	colors.GreenBold()
	fmt.Printf("\nExiting.\n\n")
	colors.Reset()
}
