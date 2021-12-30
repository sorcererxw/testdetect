package testdetect

import (
	"os"
	"path"
)

// IsTesting returns true if the process is being run as part of go test.
func IsTesting() bool {
	args := os.Args
	if len(args) == 0 {
		return false
	}
	return path.Ext(args[0]) == ".test"
}
