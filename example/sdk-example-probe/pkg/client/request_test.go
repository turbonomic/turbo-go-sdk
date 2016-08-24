package client

import (
	"net/url"
	"testing"
)

func TestParam(t *testing.T) {
	table := []struct {
		name      string
		testVal   string
		expectStr string
	}{
		{"foo", "31415", "http://localhost?foo=31415"},
		{"bar", "42", "http://localhost?bar=42"},
		{"baz", "0", "http://localhost?baz=0"},
	}

	for _, item := range table {
		u, _ := url.Parse("http://localhost")
		r := NewRequest("GET", u, "").Param(item.name, item.testVal)
		if e, a := item.expectStr, r.URL().String(); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
	}
}
