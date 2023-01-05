package types

import "errors"

type Resp struct {
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func (r *Resp) CheckErrno() error {
	if r.Errno != 0 {
		return errors.New(r.Errmsg)
	} else {
		return nil
	}
}

type KlineResp struct {
	Resp
	Result Kline `json:"result"`
}

type Kline struct {
	Ts     int64        `json:"ts"`
	Symbol string       `json:"symbol"`
	Period string       `json:"period"`
	Data   []*KlineData `json:"data"`
}

type KlineData struct {
	ID     int    `json:"id"`
	Amount string `json:"amount"`
	Count  int    `json:"count"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Vol    string `json:"vol"`
}
