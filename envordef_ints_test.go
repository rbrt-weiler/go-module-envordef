package envordef

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("INTVAL_NEG", "-10")
	os.Setenv("INTVAL_ZERO", "0")
	os.Setenv("INTVAL_POS", "10")
}

func TestIntVal(t *testing.T) {
	var testVal int

	if testVal = IntVal("INTVAL_NEG", 5); testVal != -10 {
		t.Errorf("IntVal converted %s to %d", os.Getenv("INTVAL_NEG"), testVal)
	}
	if testVal = IntVal("INTVAL_ZERO", 55); testVal != 0 {
		t.Errorf("IntVal converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = IntVal("INTVAL_POS", -5); testVal != 10 {
		t.Errorf("IntVal converted %s to %d", os.Getenv("INTVAL_POS"), testVal)
	}
}

func TestInt8Val(t *testing.T) {
	var testVal int8

	if testVal = Int8Val("INTVAL_NEG", 5); testVal != -10 {
		t.Errorf("Int8Val converted %s to %d", os.Getenv("INTVAL_NEG"), testVal)
	}
	if testVal = Int8Val("INTVAL_ZERO", 55); testVal != 0 {
		t.Errorf("Int8Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Int8Val("INTVAL_POS", -5); testVal != 10 {
		t.Errorf("Int8Val converted %s to %d", os.Getenv("INTVAL_POS"), testVal)
	}
}

func TestInt16Val(t *testing.T) {
	var testVal int16

	if testVal = Int16Val("INTVAL_NEG", 5); testVal != -10 {
		t.Errorf("Int16Val converted %s to %d", os.Getenv("INTVAL_NEG"), testVal)
	}
	if testVal = Int16Val("INTVAL_ZERO", 555); testVal != 0 {
		t.Errorf("Int16Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Int16Val("INTVAL_POS", -5); testVal != 10 {
		t.Errorf("Int16Val converted %s to %d", os.Getenv("INTVAL_POS"), testVal)
	}
}

func TestInt32Val(t *testing.T) {
	var testVal int32

	if testVal = Int32Val("INTVAL_NEG", 5); testVal != -10 {
		t.Errorf("Int32Val converted %s to %d", os.Getenv("INTVAL_NEG"), testVal)
	}
	if testVal = Int32Val("INTVAL_ZERO", 555); testVal != 0 {
		t.Errorf("Int32Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Int32Val("INTVAL_POS", -5); testVal != 10 {
		t.Errorf("Int32Val converted %s to %d", os.Getenv("INTVAL_POS"), testVal)
	}
}

func TestInt64Val(t *testing.T) {
	var testVal int64

	if testVal = Int64Val("INTVAL_NEG", 5); testVal != -10 {
		t.Errorf("Int64Val converted %s to %d", os.Getenv("INTVAL_NEG"), testVal)
	}
	if testVal = Int64Val("INTVAL_ZERO", 555); testVal != 0 {
		t.Errorf("Int64Val converted %s to %d", os.Getenv("INTVAL_ZERO"), testVal)
	}
	if testVal = Int64Val("INTVAL_POS", -5); testVal != 10 {
		t.Errorf("Int64Val converted %s to %d", os.Getenv("INTVAL_POS"), testVal)
	}
}
