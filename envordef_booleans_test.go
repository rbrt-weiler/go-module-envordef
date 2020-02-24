package envordef

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("BOOLVAL_0", "0")
	os.Setenv("BOOLVAL_1", "1")
	os.Setenv("BOOLVAL_FALSE", "false")
	os.Setenv("BOOLVAL_TRUE", "true")
}

func TestBoolVal(t *testing.T) {
	var testVal bool

	if testVal = BoolVal("BOOLVAL_0", true); testVal != false {
		t.Errorf("BoolVal converted 0 to %t", testVal)
	}
	if testVal = BoolVal("BOOLVAL_1", false); testVal != true {
		t.Errorf("BoolVal converted 1 to %t", testVal)
	}
	if testVal = BoolVal("BOOLVAL_FALSE", true); testVal != false {
		t.Errorf("BoolVal converted false to %t", testVal)
	}
	if testVal = BoolVal("BOOLVAL_TRUE", false); testVal != true {
		t.Errorf("BoolVal converted true to %t", testVal)
	}
}
