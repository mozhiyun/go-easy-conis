package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha1Hex(t *testing.T) {
	r := Sha1Hex("1534927978_ab43c57ba172a6be125cca2f449826f9980casymbol=BTC-USDTtype=1")
	assert.Equal(t, r, "731faa3d170bb746a767cea58ae563830594e1fe")
}
