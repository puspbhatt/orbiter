package autoscaler

import (
	"testing"
)

func TestFakeProvider(t *testing.T) {
	var p Provider
	a := NewAutoscaler(p, "fgaerge", 3, 4, 1)
	if a.provider != p {
		t.Fail()
	}
}
