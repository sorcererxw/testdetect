package testdetect

import (
	"testing"
)

func TestIsTesting(t *testing.T) {
	if IsTesting() != true {
		t.Fail()
	}
}
