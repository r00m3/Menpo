package hash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"menpo/colors"
	"os"
)

func CalcSha256(name string) {
	// Open file.
	file, err := os.Open(name)
	if err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error opening %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	defer file.Close()
	// Create sha256 hash.
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		colors.RedBold()
		fmt.Printf("\n    Error creating sha256 hash for %v.\n", name)
		colors.Reset()
		fmt.Println(err)
	}
	fmt.Printf("\x1b[36m\n%v\x1b[0m sha256sum is: ", name)
	fmt.Printf("\x1b[36m%x\x1b[0m\n\n", hash.Sum(nil))
}
