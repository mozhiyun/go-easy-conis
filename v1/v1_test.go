package v1

import (
	"encoding/json"
	"testing"
)

func TestClient_MarketKline(t *testing.T) {
	client := NewClient("https://oapi.easycoins.com", "", "adadadadadadadadadadada")
	resp, err := client.MarketKline("ETH-USDT", "1min", 50)
	if err != nil {
		t.Fatal(err)
	}
	s, _ := json.Marshal(resp)
	t.Log(string(s))
}
