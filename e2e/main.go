package main

import (
	"fmt"
	"os"

	"github.com/sorcererxw/testdetect"
)

func main() {
	if testdetect.IsTesting() {
		fmt.Println("should not in test")
		os.Exit(1)
	}
}
