package envordef

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("STRINGVAL", "foobar")
}

func TestStringVal(t *testing.T) {
	var testVal string

	if testVal = StringVal("STRINGVAL", ""); testVal != "foobar" {
		t.Errorf("StringVal converted %s to %s", os.Getenv("STRINGVAL"), testVal)
	}
}
