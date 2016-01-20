package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	assert := assert.New(t)
	var a string = "Hello"
	var b string = "Pam"
	for i := 0; i < 10; {
		assert.Equal(t, a, b, "they are Equal")
		assert.Equal(a, b, "they are Equal")
		assert.Equal(a, b)
		assert.Equal(1, 2)
		assert.Exactly(a, b)
	}

}
