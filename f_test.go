package f_test

import (
	"testing"

	"github.com/jhonghee/f"
)

func TestVersion(t *testing.T) {
	expected := "F v1.1.0->G v1.1.0"
	if f.Version() != expected {
		t.Error("Expected", expected, "but got", f.Version())
	}
}
