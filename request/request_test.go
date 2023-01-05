package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_MakeSignature(t *testing.T) {
	client := NewClient("testbaseurl", "57ba172a6be125c", "ca2f449826f9980ca")
	data := make(map[string]string)
	data["symbol"] = "BTC-USDT"
	data["type"] = "1"
	signature := client.MakeSignature("1534927978_ab43c", data)
	assert.Equal(t, "731faa3d170bb746a767cea58ae563830594e1fe", signature)
}
