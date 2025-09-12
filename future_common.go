package bybit

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

// FutureCommonService :
type FutureCommonService struct {
	client *Client
}

// APIKeyInfoResult :
type APIKeyInfoResult struct {
	CreatedAt     time.Time `json:"created_at"`
	ExpiredAt     time.Time `json:"expired_at"`
	VipLevel      string    `json:"vip_level"`
	Type          string    `json:"type"`
	MktMakerLevel string    `json:"mkt_maker_level"`
	APIKey        string    `json:"api_key"`
	Note          string    `json:"note"`
	Ips           []string  `json:"ips"`
	Permissions   []string  `json:"permissions"`
	InviterID     int       `json:"inviter_id"`
	UserID        int       `json:"user_id"`
	AffiliateID   int       `json:"affiliate_id"`
	ReadOnly      bool      `json:"read_only"`
}

// APIKeyInfoResponse :
type APIKeyInfoResponse struct {
	Result         []APIKeyInfoResult `json:"result"`
	CommonResponse `json:",inline"`
}

// APIKeyInfo :
func (s *FutureCommonService) APIKeyInfo(ctx context.Context) (*APIKeyInfoResponse, error) {
	var (
		res   APIKeyInfoResponse
		query url.Values
	)

	if err := s.client.getPrivately(ctx, "/v2/private/account/api-key", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// BalanceResponse :
type BalanceResponse struct {
	Result         BalanceResult `json:"result"`
	CommonResponse `json:",inline"`
}

// BalanceResult :
type BalanceResult struct {
	Balance map[Coin]Balance
}

// UnmarshalJSON :
func (r *BalanceResult) UnmarshalJSON(data []byte) error {
	parsedData := map[string]Balance{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	r.Balance = map[Coin]Balance{}
	for coin, balanceData := range parsedData {
		r.Balance[Coin(coin)] = balanceData
	}
	return nil
}

// Balance :
type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

// Balance :
func (s *FutureCommonService) Balance(ctx context.Context, coin Coin) (*BalanceResponse, error) {
	var res BalanceResponse

	query := url.Values{}
	query.Add("coin", string(coin))
	if err := s.client.getPrivately(ctx, "/v2/private/wallet/balance", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// OrderBookResponse :
type OrderBookResponse struct {
	Result         []OrderBookResult `json:"result"`
	CommonResponse `json:",inline"`
}

// OrderBookResult :
type OrderBookResult struct {
	Symbol SymbolFuture `json:"symbol"`
	Price  string       `json:"price"`
	Side   Side         `json:"side"`
	Size   float64      `json:"size"`
}

// OrderBook :
func (s *FutureCommonService) OrderBook(ctx context.Context, symbol SymbolFuture) (*OrderBookResponse, error) {
	var res OrderBookResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPublicly(ctx, "/v2/public/orderBook/L2", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListKlineParam :
type ListKlineParam struct {
	Limit    *int         `url:"limit,omitempty"`
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`
}

// ListKlineResponse :
type ListKlineResponse struct {
	Result         []ListKlineResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListKlineResult :
type ListKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Interval string       `json:"interval"`
	Open     string       `json:"open"`
	High     string       `json:"high"`
	Low      string       `json:"low"`
	Close    string       `json:"close"`
	Volume   string       `json:"volume"`
	Turnover string       `json:"turnover"`
	OpenTime int          `json:"open_time"`
}

// ListKline :
func (s *FutureCommonService) ListKline(ctx context.Context, param ListKlineParam) (*ListKlineResponse, error) {
	var res ListKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/kline/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// TickersResponse :
type TickersResponse struct {
	Result         []TickersResult `json:"result"`
	CommonResponse `json:",inline"`
}

// TickersResult :
type TickersResult struct {
	Price1hPcnt          string        `json:"price_1h_pcnt"`
	LowPrice24h          string        `json:"low_price_24h"`
	AskPrice             string        `json:"ask_price"`
	LastPrice            string        `json:"last_price"`
	LastTickDirection    TickDirection `json:"last_tick_direction"`
	PrevPrice24h         string        `json:"prev_price_24h"`
	Price24hPcnt         string        `json:"price_24h_pcnt"`
	IndexPrice           string        `json:"index_price"`
	MarkPrice            string        `json:"mark_price"`
	PrevPrice1h          string        `json:"prev_price_1h"`
	BidPrice             string        `json:"bid_price"`
	Symbol               SymbolFuture  `json:"symbol"`
	HighPrice24h         string        `json:"high_price_24h"`
	NextFundingTime      string        `json:"next_funding_time"`
	OpenValue            string        `json:"open_value"`
	TotalTurnover        string        `json:"total_turnover"`
	Turnover24h          string        `json:"turnover_24h"`
	PredictedFundingRate string        `json:"predicted_funding_rate"`
	FundingRate          string        `json:"funding_rate"`
	Volume24h            float64       `json:"volume_24h"`
	TotalVolume          float64       `json:"total_volume"`
	OpenInterest         float64       `json:"open_interest"`
	CountdownHour        float64       `json:"countdown_hour"`
}

// Tickers :
func (s *FutureCommonService) Tickers(ctx context.Context, symbol SymbolFuture) (*TickersResponse, error) {
	var res TickersResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPublicly(ctx, "/v2/public/tickers", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// TradingRecordsParam :
type TradingRecordsParam struct {
	From   *int         `url:"from,omitempty"`
	Limit  *int         `url:"limit,omitempty"`
	Symbol SymbolFuture `url:"symbol"`
}

// TradingRecordsResponse :
type TradingRecordsResponse struct {
	Result         []TradingRecordsResult `json:"result"`
	CommonResponse `json:",inline"`
}

// TradingRecordsResult :
type TradingRecordsResult struct {
	Symbol SymbolFuture `json:"symbol"`
	Side   Side         `json:"side"`
	Time   string       `json:"time"`
	ID     float64      `json:"id"`
	Price  float64      `json:"price"`
	Qty    float64      `json:"qty"`
}

// TradingRecords :
func (s *FutureCommonService) TradingRecords(ctx context.Context, param TradingRecordsParam) (*TradingRecordsResponse, error) {
	var res TradingRecordsResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/trading-records", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SymbolsResponse :
type SymbolsResponse struct {
	Result         []SymbolsResult `json:"result"`
	CommonResponse `json:",inline"`
}

// SymbolsResult :
type SymbolsResult struct {
	PriceFilter    PriceFilter    `json:"price_filter"`
	Name           string         `json:"name"`
	BaseCurrency   string         `json:"base_currency"`
	QuoteCurrency  string         `json:"quote_currency"`
	TakerFee       string         `json:"taker_fee"`
	MakerFee       string         `json:"maker_fee"`
	LeverageFilter LeverageFilter `json:"leverage_filter"`
	LotSizeFilter  LotSizeFilter  `json:"lot_size_filter"`
	PriceScale     float64        `json:"price_scale"`
}

// LeverageFilter :
type LeverageFilter struct {
	LeverageStep string  `json:"leverage_step"`
	MinLeverage  float64 `json:"min_leverage"`
	MaxLeverage  float64 `json:"max_leverage"`
}

// PriceFilter :
type PriceFilter struct {
	MinPrice string `json:"min_price"`
	MaxPrice string `json:"max_price"`
	TickSize string `json:"tick_size"`
}

// LotSizeFilter :
type LotSizeFilter struct {
	MaxTradingQty float64 `json:"max_trading_qty"`
	MinTradingQty float64 `json:"min_trading_qty"`
	QtyStep       float64 `json:"qty_step"`
}

// Symbols :
func (s *FutureCommonService) Symbols(ctx context.Context) (*SymbolsResponse, error) {
	var res SymbolsResponse

	if err := s.client.getPublicly(ctx, "/v2/public/symbols", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// MarkPriceKlineResponse :
type MarkPriceKlineResponse struct {
	Result         []MarkPriceKlineResult `json:"result"`
	CommonResponse `json:",inline"`
}

// MarkPriceKlineResult :
type MarkPriceKlineResult struct {
	Symbol  SymbolFuture `json:"symbol"`
	Period  Period       `json:"period"`
	StartAt int          `json:"start_at"`
	Open    float64      `json:"open"`
	High    float64      `json:"high"`
	Low     float64      `json:"low"`
	Close   float64      `json:"close"`
}

// MarkPriceKlineParam :
type MarkPriceKlineParam struct {
	Limit    *int         `url:"limit,omitempty"`
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`
}

// MarkPriceKline :
func (s *FutureCommonService) MarkPriceKline(ctx context.Context, param MarkPriceKlineParam) (*MarkPriceKlineResponse, error) {
	var res MarkPriceKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/mark-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// IndexPriceKlineResponse :
type IndexPriceKlineResponse struct {
	Result         []IndexPriceKlineResult `json:"result"`
	CommonResponse `json:",inline"`
}

// IndexPriceKlineResult :
type IndexPriceKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Period   Period       `json:"period"`
	Open     string       `json:"open"`
	High     string       `json:"high"`
	Low      string       `json:"low"`
	Close    string       `json:"close"`
	OpenTime int          `json:"open_time"`
}

// IndexPriceKlineParam :
type IndexPriceKlineParam struct {
	Limit    *int         `url:"limit,omitempty"`
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`
}

// IndexPriceKline :
func (s *FutureCommonService) IndexPriceKline(ctx context.Context, param IndexPriceKlineParam) (*IndexPriceKlineResponse, error) {
	var res IndexPriceKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/index-price-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// OpenInterestResponse :
type OpenInterestResponse struct {
	Result         []OpenInterestResult `json:"result"`
	CommonResponse `json:",inline"`
}

// OpenInterestResult :
type OpenInterestResult struct {
	Symbol       SymbolFuture `json:"symbol"`
	OpenInterest float64      `json:"open_interest"`
	Timestamp    int          `json:"timestamp"`
}

// OpenInterestParam :
type OpenInterestParam struct {
	Limit  *int         `url:"limit,omitempty"`
	Symbol SymbolFuture `url:"symbol"`
	Period Period       `url:"period"`
}

// OpenInterest :
func (s *FutureCommonService) OpenInterest(ctx context.Context, param OpenInterestParam) (*OpenInterestResponse, error) {
	var res OpenInterestResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/open-interest", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// BigDealResponse :
type BigDealResponse struct {
	Result         []BigDealResult `json:"result"`
	CommonResponse `json:",inline"`
}

// BigDealResult :
type BigDealResult struct {
	Symbol    SymbolFuture `json:"symbol"`
	Side      Side         `json:"side"`
	Timestamp int          `json:"timestamp"`
	Value     float64      `json:"value"`
}

// BigDealParam :
type BigDealParam struct {
	Limit  *int         `url:"limit,omitempty"`
	Symbol SymbolFuture `url:"symbol"`
}

// BigDeal :
func (s *FutureCommonService) BigDeal(ctx context.Context, param BigDealParam) (*BigDealResponse, error) {
	var res BigDealResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/big-deal", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// AccountRatioResponse :
type AccountRatioResponse struct {
	Result         []AccountRatioResult `json:"result"`
	CommonResponse `json:",inline"`
}

// AccountRatioResult :
type AccountRatioResult struct {
	Symbol    SymbolFuture `json:"symbol"`
	BuyRatio  float64      `json:"buy_ratio"`
	SellRatio float64      `json:"sell_ratio"`
	Timestamp int          `json:"timestamp"`
}

// AccountRatioParam :
type AccountRatioParam struct {
	Limit  *int         `url:"limit,omitempty"`
	Symbol SymbolFuture `url:"symbol"`
	Period Period       `url:"period"`
}

// AccountRatio :
func (s *FutureCommonService) AccountRatio(ctx context.Context, param AccountRatioParam) (*AccountRatioResponse, error) {
	var res AccountRatioResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/account-ratio", queryString, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
