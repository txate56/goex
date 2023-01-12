package goex

import (
	"github.com/nntaoli-project/goex/v2/model"
)

// IPubRest 公共接口，不需要授权调用
type IPubRest interface {
	GetName() string //获取交易所名字/域名
	GetDepth(pair model.CurrencyPair, limit int, opt ...model.OptionParameter) (depth *model.Depth, responseBody []byte, err error)
	GetTicker(pair model.CurrencyPair, opt ...model.OptionParameter) (ticker *model.Ticker, responseBody []byte, err error)
	GetKline(pair model.CurrencyPair, period model.KlinePeriod, opt ...model.OptionParameter) (klines []model.Kline, responseBody []byte, err error)
}

// IPrvRest 私有接口，需要授权调用
type IPrvRest interface {
	//CreateOrder
	//@returns
	//  order        包含订单ID信息
	//  responseBody 交易所接口返回的原始字节数据
	//  err          错误
	CreateOrder(pair model.CurrencyPair, qty, price float64, side model.OrderSide, orderTy model.OrderType, opt ...model.OptionParameter) (order *model.Order, responseBody []byte, err error)
	GetOrderInfo(pair model.CurrencyPair, id string, opt ...model.OptionParameter) (order *model.Order, responseBody []byte, err error)
	GetPendingOrders(pair model.CurrencyPair, opt ...model.OptionParameter) (orders []model.Order, responseBody []byte, err error)
	GetHistoryOrders(pair model.CurrencyPair, opt ...model.OptionParameter) (orders []model.Order, responseBody []byte, err error)
	CancelOrder(pair model.CurrencyPair, id string, opt ...model.OptionParameter) (responseBody []byte, err error)
}

type ISpotPrvRest interface {
	IPrvRest
	GetAccount(coin string) (map[string]model.Account, []byte, error)
}

// IFuturesPrvRest 合约接口
type IFuturesPrvRest interface {
	IPrvRest
	GetFuturesAccount(coin string) (acc map[string]model.FuturesAccount, responseBody []byte, err error)
	//GetPositions 获取持仓数据
	//@returns
	//	positions    仓位数据
	//  responseBody 交易所接口返回的原始字节数据
	//  err          错误
	GetPositions(pair model.CurrencyPair, opts ...model.OptionParameter) (positions []model.FuturesPosition, responseBody []byte, err error)
}
