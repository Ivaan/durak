package durakRules

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	c1 := CountCards(0xFF00000000000000)
	if c1 != 8 {
		t.Error("c1", c1, "should be 8")
	}
	c2 := CountCards(0xFF000000000000F1)
	if c2 != 13 {
		t.Error("c2", c2, "should be 13")
	}
}
