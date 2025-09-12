package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// SpotV1ServiceI :
type SpotV1ServiceI interface {
	// Market Data Endpoints
	SpotSymbols(context.Context) (*SpotSymbolsResponse, error)
	SpotQuoteDepth(context.Context, SpotQuoteDepthParam) (*SpotQuoteDepthResponse, error)
	SpotQuoteDepthMerged(context.Context, SpotQuoteDepthMergedParam) (*SpotQuoteDepthMergedResponse, error)
	SpotQuoteTrades(context.Context, SpotQuoteTradesParam) (*SpotQuoteTradesResponse, error)
	SpotQuoteKline(context.Context, SpotQuoteKlineParam) (*SpotQuoteKlineResponse, error)
	SpotQuoteTicker24hr(context.Context, SpotQuoteTicker24hrParam) (*SpotQuoteTicker24hrResponse, error)
	SpotQuoteTickerPrice(context.Context, SpotQuoteTickerPriceParam) (*SpotQuoteTickerPriceResponse, error)
	SpotQuoteTickerBookTicker(context.Context, SpotQuoteTickerBookTickerParam) (*SpotQuoteTickerBookTickerResponse, error)

	// Account Data Endpoints
	SpotPostOrder(context.Context, SpotPostOrderParam) (*SpotPostOrderResponse, error)
	SpotGetOrder(context.Context, SpotGetOrderParam) (*SpotGetOrderResponse, error)
	SpotDeleteOrder(context.Context, SpotDeleteOrderParam) (*SpotDeleteOrderResponse, error)
	SpotDeleteOrderFast(context.Context, SpotDeleteOrderFastParam) (*SpotDeleteOrderFastResponse, error)
	SpotOrderBatchCancel(context.Context, SpotOrderBatchCancelParam) (*SpotOrderBatchCancelResponse, error)
	SpotOrderBatchFastCancel(context.Context, SpotOrderBatchFastCancelParam) (*SpotOrderBatchFastCancelResponse, error)
	SpotOrderBatchCancelByIDs(ctx context.Context, orderIDs []string) (*SpotOrderBatchCancelByIDsResponse, error)
	SpotOpenOrders(context.Context, SpotOpenOrdersParam) (*SpotOpenOrdersResponse, error)

	// Wallet Data Endpoints
	SpotGetWalletBalance(context.Context) (*SpotGetWalletBalanceResponse, error)
}

// SpotV1Service :
type SpotV1Service struct {
	client *Client
}

// SpotSymbolsResponse :
type SpotSymbolsResponse struct {
	Result         []SpotSymbolsResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotSymbolsResult :
type SpotSymbolsResult struct {
	Name              string `json:"name"`
	Alias             string `json:"alias"`
	BaseCurrency      string `json:"baseCurrency"`
	QuoteCurrency     string `json:"quoteCurrency"`
	BasePrecision     string `json:"basePrecision"`
	QuotePrecision    string `json:"quotePrecision"`
	MinTradeQuantity  string `json:"minTradeQuantity"`
	MinTradeAmount    string `json:"minTradeAmount"`
	MinPricePrecision string `json:"minPricePrecision"`
	MaxTradeQuantity  string `json:"maxTradeQuantity"`
	MaxTradeAmount    string `json:"maxTradeAmount"`
	Category          int    `json:"category"`
}

// SpotSymbols :
func (s *SpotV1Service) SpotSymbols(ctx context.Context) (*SpotSymbolsResponse, error) {
	var res SpotSymbolsResponse

	if err := s.client.getPublicly(ctx, "/spot/v1/symbols", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteDepthParam :
type SpotQuoteDepthParam struct {
	Limit  *int       `url:"limit,omitempty"`
	Symbol SymbolSpot `url:"symbol"`
}

// SpotQuoteDepthResponse :
type SpotQuoteDepthResponse struct {
	Result         SpotQuoteDepthResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteDepthResult :
type SpotQuoteDepthResult struct {
	Bids SpotQuoteDepthBidsAsks `json:"bids"`
	Asks SpotQuoteDepthBidsAsks `json:"asks"`
	Time int                    `json:"time"`
}

// SpotQuoteDepthBidsAsks :
type SpotQuoteDepthBidsAsks []SpotQuoteDepthBidAsk

// UnmarshalJSON :
func (r *SpotQuoteDepthBidsAsks) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := SpotQuoteDepthBidsAsks{}
	for _, item := range parsedData {
		item := item
		if len(item) != 2 {
			return errors.New("so far len(item) must be 2, please check it on documents")
		}
		items = append(items, SpotQuoteDepthBidAsk{
			Price:    item[0],
			Quantity: item[1],
		})
	}
	*r = items
	return nil
}

// SpotQuoteDepthBidAsk :
type SpotQuoteDepthBidAsk struct {
	Price    string
	Quantity string
}

// SpotQuoteDepth :
func (s *SpotV1Service) SpotQuoteDepth(ctx context.Context, param SpotQuoteDepthParam) (*SpotQuoteDepthResponse, error) {
	var res SpotQuoteDepthResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/depth", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteDepthMergedParam :
type SpotQuoteDepthMergedParam struct {
	Scale  *int       `url:"scale,omitempty"`
	Limit  *int       `url:"limit,omitempty"`
	Symbol SymbolSpot `url:"symbol"`
}

// SpotQuoteDepthMergedResponse :
type SpotQuoteDepthMergedResponse struct {
	Result         SpotQuoteDepthMergedResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteDepthMergedResult :
type SpotQuoteDepthMergedResult struct {
	Bids SpotQuoteDepthBidsAsks `json:"bids"`
	Asks SpotQuoteDepthBidsAsks `json:"asks"`
	Time int                    `json:"time"`
}

// SpotQuoteDepthMerged :
func (s *SpotV1Service) SpotQuoteDepthMerged(ctx context.Context, param SpotQuoteDepthMergedParam) (*SpotQuoteDepthMergedResponse, error) {
	var res SpotQuoteDepthMergedResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/depth/merged", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTradesParam :
type SpotQuoteTradesParam struct {
	Limit  *int       `url:"limit,omitempty"`
	Symbol SymbolSpot `url:"symbol"`
}

// SpotQuoteTradesResponse :
type SpotQuoteTradesResponse struct {
	Result         []SpotQuoteTradesResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteTradesResult :
type SpotQuoteTradesResult struct {
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	Time         int    `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
}

// SpotQuoteTrades :
func (s *SpotV1Service) SpotQuoteTrades(ctx context.Context, param SpotQuoteTradesParam) (*SpotQuoteTradesResponse, error) {
	var res SpotQuoteTradesResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/trades", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteKlineParam :
type SpotQuoteKlineParam struct {
	Limit     *int       `url:"limit,omitempty"`
	StartTime *int       `url:"startTime,omitempty"`
	EndTime   *int       `url:"endTime,omitempty"`
	Symbol    SymbolSpot `url:"symbol"`
	Interval  Interval   `url:"interval"`
}

// SpotQuoteKlineResponse :
type SpotQuoteKlineResponse struct {
	Result         []SpotQuoteKlineResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteKlineResult :
type SpotQuoteKlineResult struct {
	SpotQuoteKline SpotQuoteKline
}

// UnmarshalJSON :
func (r *SpotQuoteKlineResult) UnmarshalJSON(data []byte) error {
	parsedData := []interface{}{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	if len(parsedData) != 11 {
		return errors.New("so far len(items) must be 11, please check it on documents")
	}
	r.SpotQuoteKline = SpotQuoteKline{
		StartTime:        int(parsedData[0].(float64)),
		Open:             parsedData[1].(string),
		High:             parsedData[2].(string),
		Low:              parsedData[3].(string),
		Close:            parsedData[4].(string),
		Volume:           parsedData[5].(string),
		EndTime:          int(parsedData[6].(float64)),
		QuoteAssetVolume: parsedData[7].(string),
		Trades:           int(parsedData[8].(float64)),
		TakerBaseVolume:  parsedData[9].(float64),
		TakerQuoteVolume: parsedData[10].(float64),
	}
	return nil
}

// SpotQuoteKline :
type SpotQuoteKline struct {
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	QuoteAssetVolume string
	StartTime        int
	EndTime          int
	Trades           int
	TakerBaseVolume  float64
	TakerQuoteVolume float64
}

// SpotQuoteKline :
func (s *SpotV1Service) SpotQuoteKline(ctx context.Context, param SpotQuoteKlineParam) (*SpotQuoteKlineResponse, error) {
	var res SpotQuoteKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTicker24hrParam :
type SpotQuoteTicker24hrParam struct {
	Symbol *SymbolSpot `url:"symbol,omitempty"`
}

// SpotQuoteTicker24hrResponse :
type SpotQuoteTicker24hrResponse struct {
	Result         SpotQuoteTicker24hrResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteTicker24hrResult :
type SpotQuoteTicker24hrResult struct {
	Symbol       string `json:"symbol"`
	BestBidPrice string `json:"bestBidPrice"`
	BestAskPrice string `json:"bestAskPrice"`
	LastPrice    string `json:"lastPrice"`
	OpenPrice    string `json:"openPrice"`
	HighPrice    string `json:"highPrice"`
	LowPrice     string `json:"lowPrice"`
	Volume       string `json:"volume"`
	QuoteVolume  string `json:"quoteVolume"`
	Time         int    `json:"time"`
}

// SpotQuoteTicker24hr :
func (s *SpotV1Service) SpotQuoteTicker24hr(ctx context.Context, param SpotQuoteTicker24hrParam) (*SpotQuoteTicker24hrResponse, error) {
	var res SpotQuoteTicker24hrResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/ticker/24hr", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTickerPriceParam :
type SpotQuoteTickerPriceParam struct {
	Symbol *SymbolSpot `url:"symbol,omitempty"`
}

// SpotQuoteTickerPriceResponse :
type SpotQuoteTickerPriceResponse struct {
	Result         SpotQuoteTickerPriceResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteTickerPriceResult :
type SpotQuoteTickerPriceResult struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// SpotQuoteTickerPrice :
func (s *SpotV1Service) SpotQuoteTickerPrice(ctx context.Context, param SpotQuoteTickerPriceParam) (*SpotQuoteTickerPriceResponse, error) {
	var res SpotQuoteTickerPriceResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/ticker/price", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotQuoteTickerBookTickerParam :
type SpotQuoteTickerBookTickerParam struct {
	Symbol *SymbolSpot `url:"symbol,omitempty"`
}

// SpotQuoteTickerBookTickerResponse :
type SpotQuoteTickerBookTickerResponse struct {
	Result         SpotQuoteTickerBookTickerResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotQuoteTickerBookTickerResult :
type SpotQuoteTickerBookTickerResult struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
	Time     int    `json:"time"`
}

// SpotQuoteTickerBookTicker :
func (s *SpotV1Service) SpotQuoteTickerBookTicker(ctx context.Context, param SpotQuoteTickerBookTickerParam) (*SpotQuoteTickerBookTickerResponse, error) {
	var res SpotQuoteTickerBookTickerResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/spot/quote/v1/ticker/book_ticker", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotPostOrderParam :
type SpotPostOrderParam struct {
	TimeInForce *TimeInForceSpot `url:"timeInForce,omitempty"`
	Price       *float64         `url:"price,omitempty"`
	OrderLinkID *string          `url:"orderLinkId,omitempty"`
	Symbol      SymbolSpot       `url:"symbol"`
	Side        Side             `url:"side"`
	Type        OrderTypeSpot    `url:"type"`
	Qty         float64          `url:"qty"`
}

// SpotPostOrderResponse :
type SpotPostOrderResponse struct {
	Result         SpotPostOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotPostOrderResult :
type SpotPostOrderResult struct {
	OrderID      string          `json:"orderId"`
	OrderLinkID  string          `json:"orderLinkId"`
	Symbol       string          `json:"symbol"`
	TransactTime string          `json:"transactTime"`
	Price        string          `json:"price"`
	OrigQty      string          `json:"origQty"`
	Type         OrderTypeSpot   `json:"type"`
	Side         string          `json:"side"`
	Status       OrderStatusSpot `json:"status"`
	TimeInForce  TimeInForceSpot `json:"timeInForce"`
	AccountID    string          `json:"accountId"`
	SymbolName   string          `json:"symbolName"`
	ExecutedQty  string          `json:"executedQty"`
}

// SpotPostOrder :
func (s *SpotV1Service) SpotPostOrder(ctx context.Context, param SpotPostOrderParam) (*SpotPostOrderResponse, error) {
	var res SpotPostOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postForm(ctx, "/spot/v1/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotGetOrderParam :
type SpotGetOrderParam struct {
	OrderID     *string `url:"orderId,omitempty"`
	OrderLinkID *string `url:"orderLinkId,omitempty"`
}

// SpotGetOrderResponse :
type SpotGetOrderResponse struct {
	Result         SpotGetOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotGetOrderResult :
type SpotGetOrderResult struct {
	AccountID           string `json:"accountId"`
	ExchangeID          string `json:"exchangeId"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbolName"`
	OrderLinkID         string `json:"orderLinkId"`
	OrderID             string `json:"orderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
}

// SpotGetOrder :
func (s *SpotV1Service) SpotGetOrder(ctx context.Context, param SpotGetOrderParam) (*SpotGetOrderResponse, error) {
	var res SpotGetOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/spot/v1/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotDeleteOrderParam :
type SpotDeleteOrderParam struct {
	OrderID     *string `url:"orderId,omitempty"`
	OrderLinkID *string `url:"orderLinkId,omitempty"`
}

// SpotDeleteOrderResponse :
type SpotDeleteOrderResponse struct {
	Result         SpotDeleteOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotDeleteOrderResult :
type SpotDeleteOrderResult struct {
	OrderID      string `json:"orderId"`
	OrderLinkID  string `json:"orderLinkId"`
	Symbol       string `json:"symbol"`
	Status       string `json:"status"`
	AccountID    string `json:"accountId"`
	TransactTime string `json:"transactTime"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	ExecutedQty  string `json:"executedQty"`
	TimeInForce  string `json:"timeInForce"`
	Type         string `json:"type"`
	Side         string `json:"side"`
}

// SpotDeleteOrder :
func (s *SpotV1Service) SpotDeleteOrder(ctx context.Context, param SpotDeleteOrderParam) (*SpotDeleteOrderResponse, error) {
	var res SpotDeleteOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately(ctx, "/spot/v1/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotDeleteOrderFastParam struct {
	OrderID     *string    `url:"orderId,omitempty"`
	OrderLinkID *string    `url:"orderLinkId,omitempty"`
	Symbol      SymbolSpot `url:"symbolId"`
}

type SpotDeleteOrderFastResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotDeleteOrderFastResult `json:"result"`
}

type SpotDeleteOrderFastResult struct {
	IsCancelled bool `json:"isCancelled"`
}

// SpotDeleteOrderFast :
func (s *SpotV1Service) SpotDeleteOrderFast(ctx context.Context, param SpotDeleteOrderFastParam) (*SpotDeleteOrderFastResponse, error) {
	var res SpotDeleteOrderFastResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately(ctx, "/spot/v1/order/fast", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchCancelParam struct {
	Symbol SymbolSpot `url:"symbolId"`

	Side  *Side           `url:"side,omitempty"`
	Types []OrderTypeSpot `url:"orderTypes,omitempty" del:","`
}

type SpotOrderBatchCancelResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotOrderBatchCancelResult `json:"result"`
}

type SpotOrderBatchCancelResult struct {
	Success bool `json:"success"`
}

func (s *SpotV1Service) SpotOrderBatchCancel(ctx context.Context, param SpotOrderBatchCancelParam) (*SpotOrderBatchCancelResponse, error) {
	var res SpotOrderBatchCancelResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately(ctx, "/spot/order/batch-cancel", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchFastCancelParam struct {
	Symbol SymbolSpot `url:"symbolId"`

	Side  *Side           `url:"side,omitempty"`
	Types []OrderTypeSpot `url:"orderTypes,omitempty" del:","`
}

type SpotOrderBatchFastCancelResponse struct {
	CommonResponse `json:",inline"`
	Result         SpotOrderBatchFastCancelResult `json:"result"`
}

type SpotOrderBatchFastCancelResult struct {
	Success bool `json:"success"`
}

func (s *SpotV1Service) SpotOrderBatchFastCancel(ctx context.Context, param SpotOrderBatchFastCancelParam) (*SpotOrderBatchFastCancelResponse, error) {
	var res SpotOrderBatchFastCancelResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.deletePrivately(ctx, "/spot/order/batch-fast-cancel", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type SpotOrderBatchCancelByIDsResponse struct {
	Result         []SpotOrderBatchCancelByIDsResult `json:"result"`
	CommonResponse `json:",inline"`
}

type SpotOrderBatchCancelByIDsResult struct {
	OrderID string `json:"orderId"`
	Code    string `json:"code"`
}

// TODO : have bug multiple orderIds
func (s *SpotV1Service) SpotOrderBatchCancelByIDs(ctx context.Context, orderIDs []string) (*SpotOrderBatchCancelByIDsResponse, error) {
	var res SpotOrderBatchCancelByIDsResponse

	if len(orderIDs) > 100 {
		return nil, errors.New("orderIDs length must be no more than 100")
	}

	query := url.Values{}
	query.Add("orderIds", strings.Join(orderIDs, ","))
	if err := s.client.deletePrivately(ctx, "/spot/order/batch-cancel-by-ids", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotOpenOrdersParam :
type SpotOpenOrdersParam struct {
	Symbol  *SymbolSpot `url:"symbol,omitempty"`
	OrderID *string     `url:"orderId,omitempty"`
	Limit   *int        `url:"limit,omitempty"`
}

// SpotOpenOrdersResponse :
type SpotOpenOrdersResponse struct {
	Result         []SpotOpenOrdersResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotOpenOrdersResult :
type SpotOpenOrdersResult struct {
	AccountID           string `json:"accountId"`
	ExchangeID          string `json:"exchangeId"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbolName"`
	OrderLinkID         string `json:"orderLinkId"`
	OrderID             string `json:"orderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
}

// SpotOpenOrders :
func (s *SpotV1Service) SpotOpenOrders(ctx context.Context, param SpotOpenOrdersParam) (*SpotOpenOrdersResponse, error) {
	var res SpotOpenOrdersResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/spot/v1/open-orders", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SpotGetWalletBalanceResponse :
type SpotGetWalletBalanceResponse struct {
	Result         SpotGetWalletBalanceResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SpotGetWalletBalanceResult :
type SpotGetWalletBalanceResult struct {
	Balances []SpotGetWalletBalanceResultBalance `json:"balances"`
}

// SpotGetWalletBalanceResultBalance :
type SpotGetWalletBalanceResultBalance struct {
	Coin     string `json:"coin"`
	CoinID   string `json:"coinId"`
	CoinName string `json:"coinName"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}

// SpotGetWalletBalance :
func (s *SpotV1Service) SpotGetWalletBalance(ctx context.Context) (*SpotGetWalletBalanceResponse, error) {
	var res SpotGetWalletBalanceResponse

	queryString, err := query.Values(nil)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/spot/v1/account", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
