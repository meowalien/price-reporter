package protocol

import "github.com/adshao/go-binance/v2"

type CoinPriceBody struct {
	//Symbol       string               `json:"symbol"`
	WsTradeEvent binance.WsTradeEvent `json:"wsTradeEvent"`
}
