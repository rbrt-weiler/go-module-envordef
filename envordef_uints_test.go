package envordef

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("INTVAL_ZERO", "0")
	os.Setenv("INTVAL_TEN", "10")
}

func TestUintVal(t *testing.T) {
	var testVal uint

	if testVal = UintVal("INTVAL_ZERO", 10); testVal != 0 {
		t.Errorf("UintVal converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = UintVal("INTVAL_TEN", 0); testVal != 10 {
		t.Errorf("UintVal converted %s to %d", os.Getenv("INTVAL_TEN"), testVal)
	}
}

func TestUint8Val(t *testing.T) {
	var testVal uint8

	if testVal = Uint8Val("INTVAL_ZERO", 10); testVal != 0 {
		t.Errorf("Uint8Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Uint8Val("INTVAL_TEN", 0); testVal != 10 {
		t.Errorf("Uint8Val converted %s to %d", os.Getenv("INTVAL_TEN"), testVal)
	}
}

func TestUint16Val(t *testing.T) {
	var testVal uint16

	if testVal = Uint16Val("INTVAL_ZERO", 10); testVal != 0 {
		t.Errorf("Uint16Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Uint16Val("INTVAL_TEN", 0); testVal != 10 {
		t.Errorf("Uint16Val converted %s to %d", os.Getenv("INTVAL_TEN"), testVal)
	}
}

func TestUint32Val(t *testing.T) {
	var testVal uint32

	if testVal = Uint32Val("INTVAL_ZERO", 10); testVal != 0 {
		t.Errorf("Uint32Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Uint32Val("INTVAL_TEN", 0); testVal != 10 {
		t.Errorf("Uint32Val converted %s to %d", os.Getenv("INTVAL_TEN"), testVal)
	}
}

func TestUint64Val(t *testing.T) {
	var testVal uint64

	if testVal = Uint64Val("INTVAL_ZERO", 10); testVal != 0 {
		t.Errorf("Uint64Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Uint64Val("INTVAL_TEN", 0); testVal != 10 {
		t.Errorf("Uint64Val converted %s to %d", os.Getenv("INTVAL_TEN"), testVal)
	}
}
