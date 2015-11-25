package rand

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	valid := "0123456789abcdefghijklmnopqrstuvwxyz"
	for _, l := range []int{0, 1, 2, 10, 123} {
		s := String(l)
		if len(s) != l {
			t.Errorf("expected string of size %d, got %q", l, s)
		}
		for _, c := range s {
			if !strings.ContainsRune(valid, c) {
				t.Errorf("expected valid charaters, got %v", c)
			}
		}
	}
}
