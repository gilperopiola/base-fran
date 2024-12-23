package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBaseFran(t *testing.T) {
	nBaseFran := toBaseFran(50000)
	assert.Equal(t, nBaseFran, 2)
}
