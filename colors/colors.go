package colors

import "fmt"

func RedBold() {
	fmt.Printf("\x1b[1;31m")
}

func GreenBold() {
	fmt.Printf("\x1b[1;32m")
}

func Cyan() {
	fmt.Printf("\x1b[36m")
}

func Reset() {
	fmt.Printf("\x1b[0;22m")
}
