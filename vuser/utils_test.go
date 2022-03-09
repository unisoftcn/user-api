package vuser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAncestors(t *testing.T) {
	path := "CAAA:DAAA:EAAA"

	ancestors := Ancestors(path)

	assert.Equal(t, 3, len(ancestors))
	assert.Equal(t, "CAAA", ancestors[0])
	assert.Equal(t, "CAAA:DAAA", ancestors[1])
	assert.Equal(t, "CAAA:DAAA:EAAA", ancestors[2])

	path = ""
	ancestors = Ancestors(path)
	assert.Equal(t, 0, len(ancestors))

	path = "CAAA"
	ancestors = Ancestors(path)
	assert.Equal(t, 1, len(ancestors))
	assert.Equal(t, "CAAA", ancestors[0])
}
