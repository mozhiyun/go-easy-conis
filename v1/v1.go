package v1

import (
	"encoding/json"
	"github.com/mozhiyun/go-easy-conis/request"
	"github.com/mozhiyun/go-easy-conis/utils"
	"github.com/mozhiyun/go-easy-conis/v1/types"
)

type Client struct {
	rc *request.Client
}

func NewClient(baseUrl, token, secretKey string) *Client {
	return &Client{
		rc: request.NewClient(baseUrl, token, secretKey),
	}
}

func (client *Client) MarketKline(symbol, period string, size int) (*types.Kline, error) {
	requestData := make(map[string]string)
	requestData["symbol"] = symbol
	requestData["period"] = period
	if size != 0 {
		requestData["size"] = utils.Int2String(size)
	}
	respByte, err := client.rc.MakeReq("/openApi/market/kline", requestData)
	if err != nil {
		return nil, err
	}
	var resp types.KlineResp
	err = json.Unmarshal(respByte, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, resp.CheckErrno()
}
