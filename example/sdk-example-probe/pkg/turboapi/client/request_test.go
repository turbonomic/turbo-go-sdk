package client

import (
	"net/url"
	"testing"

	"github.com/vmturbo/vmturbo-go-sdk/example/sdk-example-probe/pkg/api"
)

func TestParam(t *testing.T) {
	u, _ := url.Parse("http://localhost")
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
		r := NewRequest("GET", u, "").Param(item.name, item.testVal)
		if e, a := item.expectStr, r.URL().String(); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
	}
}

func TestName(t *testing.T) {
	u, _ := url.Parse("http://localhost")
	tests := []struct {
		name      string
		expectStr string
	}{
		{"bar", u.String() + "/bar"},
		{"foo", u.String() + "/foo"},
	}
	for _, test := range tests {
		r := NewRequest("GET", u, "").Name(test.name)
		if e, a := test.expectStr, r.URL().String(); e != a {
			t.Errorf("expected %s, got %s", e, a)
		}
	}
}

func TestResource(t *testing.T) {
	u, _ := url.Parse("http://localhost")
	tests := []struct {
		resource  api.ResourceType
		expectStr string
	}{
		{api.Resource_Type_Target, u.String() + "/targets"},
		{api.Resource_Type_External_Target, u.String() + "/externaltargets"},
	}
	for _, test := range tests {
		r := NewRequest("GET", u, "").Resource(test.resource)
		if e, a := test.expectStr, r.URL().String(); e != a {
			t.Errorf("expected %s, got %s", e, a)
		}
	}
}

func TestURLInOrder(t *testing.T) {
	u, _ := url.Parse("http://localhost")
	tests := []struct {
		resource     api.ResourceType
		resourceName string
		parameters   map[string]string
		expectStr    string
	}{
		{
			resource:     api.Resource_Type_Target,
			resourceName: "foo",
			expectStr:    u.String() + "/targets/foo",
		},
		{
			resource: api.Resource_Type_External_Target,
			parameters: map[string]string{
				"foo": "12",
			},
			expectStr: u.String() + "/externaltargets?foo=12"},
	}
	for _, test := range tests {
		r := NewRequest("GET", u, "").Resource(test.resource).Name(test.resourceName)
		for key, val := range test.parameters {
			r.Param(key, val)
		}
		if e, a := test.expectStr, r.URL().String(); e != a {
			t.Errorf("expected %s, got %s", e, a)
		}
	}
}
