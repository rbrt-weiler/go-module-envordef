package envordef

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("FLOATVAL", "1.234")
}

func TestFloat32Val(t *testing.T) {
	var testVal float32

	if testVal = Float32Val("FLOATVAL", 0.0); testVal != 1.234 {
		t.Errorf("Float32Val converted 1.234 to %f", testVal)
	}
}

func TestFloat64Val(t *testing.T) {
	var testVal float64

	if testVal = Float64Val("FLOATVAL", 0.0); testVal != 1.234 {
		t.Errorf("Float64Val converted 1.234 to %f", testVal)
	}
}
