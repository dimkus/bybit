package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FutureInversePerpetualServiceI :
type FutureInversePerpetualServiceI interface {
	// Market Data Endpoints
	OrderBook(context.Context, SymbolFuture) (*OrderBookResponse, error)
	ListKline(context.Context, ListKlineParam) (*ListKlineResponse, error)
	Tickers(context.Context, SymbolFuture) (*TickersResponse, error)
	TradingRecords(context.Context, TradingRecordsParam) (*TradingRecordsResponse, error)
	Symbols(context.Context) (*SymbolsResponse, error)
	MarkPriceKline(context.Context, MarkPriceKlineParam) (*MarkPriceKlineResponse, error)
	IndexPriceKline(context.Context, IndexPriceKlineParam) (*IndexPriceKlineResponse, error)
	PremiumIndexKline(context.Context, PremiumIndexKlineParam) (*PremiumIndexKlineResponse, error)
	OpenInterest(context.Context, OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(context.Context, BigDealParam) (*BigDealResponse, error)
	AccountRatio(context.Context, AccountRatioParam) (*AccountRatioResponse, error)

	// Account Data Endpoints
	CreateOrder(context.Context, CreateOrderParam) (*CreateOrderResponse, error)
	ListOrder(context.Context, ListOrderParam) (*ListOrderResponse, error)
	CancelOrder(context.Context, CancelOrderParam) (*CancelOrderResponse, error)
	CancelAllOrder(context.Context, CancelAllOrderParam) (*CancelAllOrderResponse, error)
	QueryOrder(context.Context, QueryOrderParam) (*QueryOrderResponse, error)
	CreateStopOrder(context.Context, CreateStopOrderParam) (*CreateStopOrderResponse, error)
	ListStopOrder(context.Context, ListStopOrderParam) (*ListStopOrderResponse, error)
	CancelStopOrder(context.Context, CancelStopOrderParam) (*CancelStopOrderResponse, error)
	CancelAllStopOrder(context.Context, CancelAllStopOrderParam) (*CancelAllStopOrderResponse, error)
	QueryStopOrder(context.Context, QueryStopOrderParam) (*QueryStopOrderResponse, error)
	ListPosition(context.Context, SymbolFuture) (*ListPositionResponse, error)
	ListPositions(context.Context) (*ListPositionsResponse, error)
	TradingStop(context.Context, TradingStopParam) (*TradingStopResponse, error)
	SaveLeverage(context.Context, SaveLeverageParam) (*SaveLeverageResponse, error)
	APIKeyInfo(context.Context) (*APIKeyInfoResponse, error)

	// Wallet Data Endpoints
	Balance(context.Context, Coin) (*BalanceResponse, error)
}

// FutureInversePerpetualService :
type FutureInversePerpetualService struct {
	client *Client

	*FutureCommonService
}

// PremiumIndexKlineResponse :
type PremiumIndexKlineResponse struct {
	Result         []PremiumIndexKlineResult `json:"result"`
	CommonResponse `json:",inline"`
}

// PremiumIndexKlineResult :
type PremiumIndexKlineResult struct {
	Symbol   SymbolFuture `json:"symbol"`
	Period   Period       `json:"period"`
	Open     string       `json:"open"`
	High     string       `json:"high"`
	Low      string       `json:"low"`
	Close    string       `json:"close"`
	OpenTime int          `json:"open_time"`
}

// PremiumIndexKlineParam :
type PremiumIndexKlineParam struct {
	Limit    *int         `url:"limit,omitempty"`
	Symbol   SymbolFuture `url:"symbol"`
	Interval Interval     `url:"interval"`
	From     int64        `url:"from"`
}

// PremiumIndexKline :
func (s *FutureInversePerpetualService) PremiumIndexKline(ctx context.Context, param PremiumIndexKlineParam) (*PremiumIndexKlineResponse, error) {
	var res PremiumIndexKlineResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly(ctx, "/v2/public/premium-index-kline", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateOrderResponse :
type CreateOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateOrderResult `json:"result"`
}

// CreateOrderResult :
type CreateOrderResult struct {
	CreateOrder `json:",inline"`
}

// CreateOrder :
type CreateOrder struct {
	OrderStatus   OrderStatus  `json:"order_status"`
	Side          Side         `json:"side"`
	Symbol        SymbolFuture `json:"symbol"`
	UpdatedAt     string       `json:"updated_at"`
	OrderType     OrderType    `json:"order_type"`
	CreatedAt     string       `json:"created_at"`
	OrderLinkID   string       `json:"order_link_id"`
	TimeInForce   TimeInForce  `json:"time_in_force"`
	OrderID       string       `json:"order_id"`
	RejectReason  string       `json:"reject_reason"`
	LastExecPrice float64      `json:"last_exec_price"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  float64      `json:"cum_exec_value"`
	CumExecFee    float64      `json:"cum_exec_fee"`
	UserID        int          `json:"user_id"`
	Qty           float64      `json:"qty"`
	Price         float64      `json:"price"`
	LastExecTime  float64      `json:"last_exec_time"`
}

// CreateOrderParam :
type CreateOrderParam struct {
	Price          *float64     `json:"price,omitempty"`
	TakeProfit     *float64     `json:"take_profit,omitempty"`
	StopLoss       *float64     `json:"stop_loss,omitempty"`
	ReduceOnly     *bool        `json:"reduce_only,omitempty"`
	CloseOnTrigger *bool        `json:"close_on_trigger,omitempty"`
	OrderLinkID    *string      `json:"order_link_id,omitempty"`
	Side           Side         `json:"side"`
	Symbol         SymbolFuture `json:"symbol"`
	OrderType      OrderType    `json:"order_type"`
	TimeInForce    TimeInForce  `json:"time_in_force"`
	Qty            int          `json:"qty"`
}

// CreateOrder :
func (s *FutureInversePerpetualService) CreateOrder(ctx context.Context, param CreateOrderParam) (*CreateOrderResponse, error) {
	var res CreateOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListOrderResponse :
type ListOrderResponse struct {
	Result         ListOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListOrderResult :
type ListOrderResult struct {
	ListOrders []ListOrder `json:"data"`
}

// ListOrder :
type ListOrder struct {
	CumExecQty   string          `json:"cum_exec_qty"`
	CumExecValue string          `json:"cum_exec_value"`
	Side         Side            `json:"side"`
	OrderType    OrderType       `json:"order_type"`
	Price        string          `json:"price"`
	Qty          string          `json:"qty"`
	TimeInForce  TimeInForce     `json:"time_in_force"`
	OrderStatus  OrderStatus     `json:"order_status"`
	LeavesQty    string          `json:"leaves_qty"`
	LeavesValue  string          `json:"leaves_value"`
	Symbol       SymbolFuture    `json:"symbol"`
	CumExecFee   string          `json:"cum_exec_fee"`
	SlTriggerBy  TriggerByFuture `json:"sl_trigger_by"`
	RejectReason string          `json:"reject_reason"`
	OrderLinkID  string          `json:"order_link_id"`
	CreatedAt    string          `json:"created_at"`
	OrderID      string          `json:"order_id"`
	TakeProfit   string          `json:"take_profit"`
	StopLoss     string          `json:"stop_loss"`
	TpTriggerBy  TriggerByFuture `json:"tp_trigger_by"`
	UserID       int             `json:"user_id"`
}

// ListOrderParam :
type ListOrderParam struct {
	OrderStatus *OrderStatus `url:"order_status,omitempty"`
	Direction   *Direction   `url:"direction,omitempty"`
	Size        *int         `url:"size,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
	Symbol      SymbolFuture `url:"symbol"`
}

// ListOrder :
func (s *FutureInversePerpetualService) ListOrder(ctx context.Context, param ListOrderParam) (*ListOrderResponse, error) {
	var res ListOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/v2/private/order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListPositionResponse :
type ListPositionResponse struct {
	CommonResponse `json:",inline"`
	Result         ListPositionResult `json:"result"`
}

// ListPositionResult :
type ListPositionResult struct {
	WalletBalance       string       `json:"wallet_balance"`
	TakeProfit          string       `json:"take_profit"`
	UpdatedAt           string       `json:"updated_at"`
	Symbol              SymbolFuture `json:"symbol"`
	Side                Side         `json:"side"`
	CreatedAt           string       `json:"created_at"`
	PositionValue       string       `json:"position_value"`
	EntryPrice          string       `json:"entry_price"`
	OccFundingFee       string       `json:"occ_funding_fee"`
	CumRealisedPnl      string       `json:"cum_realised_pnl"`
	Leverage            string       `json:"leverage"`
	EffectiveLeverage   string       `json:"effective_leverage"`
	PositionMargin      string       `json:"position_margin"`
	LiqPrice            string       `json:"liq_price"`
	RealisedPnl         string       `json:"realised_pnl"`
	BustPrice           string       `json:"bust_price"`
	OrderMargin         string       `json:"order_margin"`
	OccClosingFee       string       `json:"occ_closing_fee"`
	StopLoss            string       `json:"stop_loss"`
	TrailingStop        string       `json:"trailing_stop"`
	PositionStatus      string       `json:"position_status"`
	OcCalcData          string       `json:"oc_calc_data"`
	DeleverageIndicator int          `json:"deleverage_indicator"`
	ID                  int          `json:"id"`
	UserID              int          `json:"user_id"`
	UnrealisedPnl       float64      `json:"unrealised_pnl"`
	AutoAddMargin       float64      `json:"auto_add_margin"`
	CrossSeq            float64      `json:"cross_seq"`
	PositionSeq         float64      `json:"position_seq"`
	Size                float64      `json:"size"`
	RiskID              int          `json:"risk_id"`
	IsIsolated          bool         `json:"is_isolated"`
}

// ListPosition :
func (s *FutureInversePerpetualService) ListPosition(ctx context.Context, symbol SymbolFuture) (*ListPositionResponse, error) {
	var res ListPositionResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPrivately(ctx, "/v2/private/position/list", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListPositionsResponse :
type ListPositionsResponse struct {
	Result         []ListPositionsResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListPositionsResult :
type ListPositionsResult struct {
	ListPositionResult `json:"data,inline"`
	IsValid            bool `json:"is_valid"`
}

// ListPositions :
func (s *FutureInversePerpetualService) ListPositions(ctx context.Context) (*ListPositionsResponse, error) {
	var res ListPositionsResponse

	if err := s.client.getPrivately(ctx, "/v2/private/position/list", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// TradingStopResponse :
type TradingStopResponse struct {
	CommonResponse `json:",inline"`
	Result         TradingStopResult `json:"result"`
}

// TradingStopResult :
type TradingStopResult struct {
	ExtFields           map[string]interface{} `json:"ext_fields"`
	PositionStatus      string                 `json:"position_status"`
	Symbol              SymbolFuture           `json:"symbol"`
	Side                Side                   `json:"side"`
	UpdatedAt           string                 `json:"updated_at"`
	CreatedAt           string                 `json:"created_at"`
	OcCalcData          string                 `json:"oc_calc_data"`
	OccFundingFee       float64                `json:"occ_funding_fee"`
	RiskID              int                    `json:"risk_id"`
	Leverage            float64                `json:"leverage"`
	PositionMargin      float64                `json:"position_margin"`
	LiqPrice            float64                `json:"liq_price"`
	BustPrice           float64                `json:"bust_price"`
	OccClosingFee       float64                `json:"occ_closing_fee"`
	ID                  int                    `json:"id"`
	TakeProfit          float64                `json:"take_profit"`
	StopLoss            float64                `json:"stop_loss"`
	TrailingStop        float64                `json:"trailing_stop"`
	AutoAddMargin       float64                `json:"auto_add_margin"`
	DeleverageIndicator int                    `json:"deleverage_indicator"`
	EntryPrice          float64                `json:"entry_price"`
	OrderMargin         float64                `json:"order_margin"`
	WalletBalance       float64                `json:"wallet_balance"`
	RealisedPnl         float64                `json:"realised_pnl"`
	CumRealisedPnl      float64                `json:"cum_realised_pnl"`
	CumCommission       float64                `json:"cum_commission"`
	CrossSeq            float64                `json:"cross_seq"`
	PositionSeq         float64                `json:"position_seq"`
	PositionValue       float64                `json:"position_value"`
	Size                float64                `json:"size"`
	UserID              int                    `json:"user_id"`
}

// TradingStopParam :
type TradingStopParam struct {
	TakeProfit        *float64         `json:"take_profit,omitempty"`
	StopLoss          *float64         `json:"stop_loss,omitempty"`
	TrailingStop      *float64         `json:"trailing_stop,omitempty"`
	TpTriggerBy       *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	SlTriggerBy       *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	NewTrailingActive *float64         `json:"new_trailing_active,omitempty"`
	SlSize            *float64         `json:"sl_size,omitempty"`
	TpSize            *float64         `json:"tp_size,omitempty"`
	Symbol            SymbolFuture     `json:"symbol"`
}

// TradingStop :
func (s *FutureInversePerpetualService) TradingStop(ctx context.Context, param TradingStopParam) (*TradingStopResponse, error) {
	var res TradingStopResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for TradingStopParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/position/trading-stop", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelOrderResponse :
type CancelOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelOrderResult `json:"result"`
}

// CancelOrderResult :
type CancelOrderResult struct {
	CancelOrder `json:",inline"`
}

// CancelOrder :
// so far, same as CreateOrder
type CancelOrder struct {
	OrderStatus   OrderStatus  `json:"order_status"`
	Side          Side         `json:"side"`
	Symbol        SymbolFuture `json:"symbol"`
	UpdatedAt     string       `json:"updated_at"`
	OrderType     OrderType    `json:"order_type"`
	CreatedAt     string       `json:"created_at"`
	OrderLinkID   string       `json:"order_link_id"`
	TimeInForce   TimeInForce  `json:"time_in_force"`
	OrderID       string       `json:"order_id"`
	RejectReason  string       `json:"reject_reason"`
	LastExecPrice float64      `json:"last_exec_price"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  float64      `json:"cum_exec_value"`
	CumExecFee    float64      `json:"cum_exec_fee"`
	UserID        int          `json:"user_id"`
	Qty           float64      `json:"qty"`
	Price         float64      `json:"price"`
	LastExecTime  float64      `json:"last_exec_time"`
}

// CancelOrderParam :
type CancelOrderParam struct {
	OrderID     *string      `json:"order_id,omitempty"`
	OrderLinkID *string      `json:"order_link_id,omitempty"`
	Symbol      SymbolFuture `json:"symbol"`
}

// CancelOrder :
func (s *FutureInversePerpetualService) CancelOrder(ctx context.Context, param CancelOrderParam) (*CancelOrderResponse, error) {
	var res CancelOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllOrderResponse :
type CancelAllOrderResponse struct {
	Result         []CancelAllOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CancelAllOrderResult :
type CancelAllOrderResult struct {
	TimeInForce TimeInForce  `json:"time_in_force"`
	CancelType  string       `json:"cancel_type"`
	CrossStatus string       `json:"cross_status"`
	Symbol      SymbolFuture `json:"symbol"`
	Side        Side         `json:"side"`
	OrderType   OrderType    `json:"order_type"`
	Price       string       `json:"price"`
	UpdatedAt   string       `json:"updated_at"`
	OrderLinkID string       `json:"order_link_id"`
	ClOrdID     string       `json:"clOrdID"`
	CreateType  string       `json:"create_type"`
	OrderStatus OrderStatus  `json:"order_status"`
	CreatedAt   string       `json:"created_at"`
	LeavesValue string       `json:"leaves_value"`
	LeavesQty   float64      `json:"leaves_qty"`
	Qty         float64      `json:"qty"`
	UserID      int          `json:"user_id"`
	CrossSeq    int          `json:"cross_seq"`
}

// CancelAllOrderParam :
type CancelAllOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllOrder :
func (s *FutureInversePerpetualService) CancelAllOrder(ctx context.Context, param CancelAllOrderParam) (*CancelAllOrderResponse, error) {
	var res CancelAllOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryOrderResponse :
type QueryOrderResponse struct {
	Result         []QueryOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// QueryOrderResult :
type QueryOrderResult struct {
	ExtFields    map[string]interface{} `json:"ext_fields"`
	OrderLinkID  string                 `json:"order_link_id"`
	TakeProfit   string                 `json:"take_profit"`
	Side         Side                   `json:"side"`
	OrderType    OrderType              `json:"order_type"`
	Price        string                 `json:"price"`
	SlTriggerBy  TriggerByFuture        `json:"sl_trigger_by"`
	TimeInForce  TimeInForce            `json:"time_in_force"`
	OrderStatus  OrderStatus            `json:"order_status"`
	TpTriggerBy  TriggerByFuture        `json:"tp_trigger_by"`
	LastExecTime string                 `json:"last_exec_time"`
	StopLoss     string                 `json:"stop_loss"`
	LeavesValue  string                 `json:"leaves_value"`
	Symbol       SymbolFuture           `json:"symbol"`
	CumExecFee   string                 `json:"cum_exec_fee"`
	CumExecValue string                 `json:"cum_exec_value"`
	RejectReason string                 `json:"reject_reason"`
	CancelType   string                 `json:"cancel_type"`
	OrderID      string                 `json:"order_id"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	UserID       int                    `json:"user_id"`
	CumExecQty   int                    `json:"cum_exec_qty"`
	LeavesQty    int                    `json:"leaves_qty"`
	PositionIdx  int                    `json:"position_idx"`
	Qty          float64                `json:"qty"`
}

// QueryOrderParam :
type QueryOrderParam struct {
	OrderID     *string      `url:"order_id,omitempty"`
	OrderLinkID *string      `url:"order_link_id,omitempty"`
	Symbol      SymbolFuture `url:"symbol"`
}

// QueryOrder :
func (s *FutureInversePerpetualService) QueryOrder(ctx context.Context, param QueryOrderParam) (*QueryOrderResponse, error) {
	var res QueryOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/v2/private/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateStopOrderResponse :
type CreateStopOrderResponse struct {
	Result         CreateStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CreateStopOrderResult :
type CreateStopOrderResult struct {
	StopPx       string          `json:"stop_px"`
	RejectReason string          `json:"reject_reason"`
	Side         Side            `json:"side"`
	OrderType    OrderType       `json:"order_type"`
	Price        string          `json:"price"`
	Qty          string          `json:"qty"`
	TimeInForce  TimeInForce     `json:"time_in_force"`
	Remark       string          `json:"remark"`
	LeavesQty    string          `json:"leaves_qty"`
	StopOrderID  string          `json:"stop_order_id"`
	Symbol       SymbolFuture    `json:"symbol"`
	StopLoss     string          `json:"stop_loss"`
	LeavesValue  string          `json:"leaves_value"`
	OrderLinkID  string          `json:"order_link_id"`
	TriggerBy    TriggerByFuture `json:"trigger_by"`
	BasePrice    string          `json:"base_price"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
	TpTriggerBy  TriggerByFuture `json:"tp_trigger_by"`
	SlTriggerBy  TriggerByFuture `json:"sl_trigger_by"`
	TakeProfit   string          `json:"take_profit"`
	UserID       int             `json:"user_id"`
}

// CreateStopOrderParam :
type CreateStopOrderParam struct {
	TriggerBy      *TriggerByFuture `json:"trigger_by,omitempty"`
	SlTriggerBy    *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	TpTriggerBy    *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	StopLoss       *float64         `json:"stop_loss,omitempty"`
	TakeProfit     *float64         `json:"take_profit,omitempty"`
	OrderLinkID    *string          `json:"order_link_id,omitempty"`
	CloseOnTrigger *bool            `json:"close_on_trigger,omitempty"`
	Price          *float64         `json:"price,omitempty"`
	TimeInForce    TimeInForce      `json:"time_in_force"`
	Side           Side             `json:"side"`
	OrderType      OrderType        `json:"order_type"`
	Symbol         SymbolFuture     `json:"symbol"`
	StopPx         float64          `json:"stop_px"`
	BasePrice      float64          `json:"base_price"`
	Qty            int              `json:"qty"`
}

// CreateStopOrder :
func (s *FutureInversePerpetualService) CreateStopOrder(ctx context.Context, param CreateStopOrderParam) (*CreateStopOrderResponse, error) {
	var res CreateStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateStopOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/stop-order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListStopOrderResponse :
type ListStopOrderResponse struct {
	Result         ListStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListStopOrderResult :
type ListStopOrderResult struct {
	ListStopOrders []ListStopOrder `json:"data"`
}

// ListStopOrder :
type ListStopOrder struct {
	TriggerBy       TriggerByFuture     `json:"trigger_by"`
	BasePrice       string              `json:"base_price"`
	StopOrderStatus OrderStatus         `json:"stop_order_status"`
	Symbol          SymbolFuture        `json:"symbol"`
	Side            Side                `json:"side"`
	OrderType       OrderType           `json:"order_type"`
	Price           string              `json:"price"`
	Qty             string              `json:"qty"`
	TimeInForce     TimeInForce         `json:"time_in_force"`
	OrderLinkID     string              `json:"order_link_id"`
	Cursor          string              `json:"cursor"`
	SlTriggerBy     TriggerByFuture     `json:"sl_trigger_by"`
	StopOrderType   StopOrderTypeFuture `json:"stop_order_type"`
	CreatedAt       string              `json:"created_at"`
	UpdatedAt       string              `json:"updated_at"`
	StopPx          string              `json:"stop_px"`
	StopOrderID     string              `json:"stop_order_id"`
	TakeProfit      string              `json:"take_profit"`
	StopLoss        string              `json:"stop_loss"`
	TpTriggerBy     TriggerByFuture     `json:"tp_trigger_by"`
	UserID          int                 `json:"user_id"`
	PositionIdx     int                 `json:"position_idx"`
}

// ListStopOrderParam :
type ListStopOrderParam struct {
	StopOrderStatus *OrderStatus `url:"stop_order_status,omitempty"`
	Direction       *Direction   `url:"direction,omitempty"`
	Limit           *int         `url:"limit,omitempty"`
	Cursor          *string      `url:"cursor,omitempty"`
	Symbol          SymbolFuture `url:"symbol"`
}

// ListStopOrder :
func (s *FutureInversePerpetualService) ListStopOrder(ctx context.Context, param ListStopOrderParam) (*ListStopOrderResponse, error) {
	var res ListStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/v2/private/stop-order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelStopOrderResponse :
type CancelStopOrderResponse struct {
	Result         CancelStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CancelStopOrderResult :
type CancelStopOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

// CancelStopOrderParam :
type CancelStopOrderParam struct {
	StopOrderID *string      `json:"stop_order_id,omitempty"`
	OrderLinkID *string      `json:"order_link_id,omitempty"`
	Symbol      SymbolFuture `json:"symbol"`
}

// CancelStopOrder :
func (s *FutureInversePerpetualService) CancelStopOrder(ctx context.Context, param CancelStopOrderParam) (*CancelStopOrderResponse, error) {
	var res CancelStopOrderResponse

	if param.StopOrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either StopOrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelStopOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/stop-order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllStopOrderResponse :
type CancelAllStopOrderResponse struct {
	Result         []CancelAllStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CancelAllStopOrderResult :
type CancelAllStopOrderResult struct {
	CancelType        string              `json:"cancel_type"`
	OrderStatus       OrderStatus         `json:"order_status"`
	ExpectedDirection string              `json:"expected_direction"`
	Symbol            SymbolFuture        `json:"symbol"`
	Side              Side                `json:"side"`
	OrderType         OrderType           `json:"order_type"`
	Price             string              `json:"price"`
	BasePrice         string              `json:"base_price"`
	TimeInForce       TimeInForce         `json:"time_in_force"`
	TriggerBy         TriggerByFuture     `json:"trigger_by"`
	OrderLinkID       string              `json:"order_link_id"`
	ClOrdID           string              `json:"clOrdID"`
	CreateType        string              `json:"create_type"`
	LeavesValue       string              `json:"leaves_value"`
	CreatedAt         string              `json:"created_at"`
	UpdatedAt         string              `json:"updated_at"`
	CrossStatus       string              `json:"cross_status"`
	StopOrderType     StopOrderTypeFuture `json:"stop_order_type"`
	CrossSeq          int                 `json:"cross_seq"`
	LeavesQty         float64             `json:"leaves_qty"`
	Qty               float64             `json:"qty"`
	UserID            int                 `json:"user_id"`
}

// CancelAllStopOrderParam :
type CancelAllStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllStopOrder :
func (s *FutureInversePerpetualService) CancelAllStopOrder(ctx context.Context, param CancelAllStopOrderParam) (*CancelAllStopOrderResponse, error) {
	var res CancelAllStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllStopOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/stop-order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryStopOrderResponse :
type QueryStopOrderResponse struct {
	Result         []QueryStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// QueryStopOrderResult :
type QueryStopOrderResult struct {
	ExtFields       map[string]interface{} `json:"ext_fields"`
	StopOrderStatus OrderStatus            `json:"stop_order_status"`
	StopLoss        string                 `json:"stop_loss"`
	Side            Side                   `json:"side"`
	OrderType       OrderType              `json:"order_type"`
	Price           string                 `json:"price"`
	TpTriggerBy     TriggerByFuture        `json:"tp_trigger_by"`
	StopPx          string                 `json:"stop_px"`
	BasePrice       string                 `json:"base_price"`
	TimeInForce     TimeInForce            `json:"time_in_force"`
	TriggerBy       TriggerByFuture        `json:"trigger_by"`
	Symbol          SymbolFuture           `json:"symbol"`
	LeavesValue     string                 `json:"leaves_value"`
	SlTriggerBy     TriggerByFuture        `json:"sl_trigger_by"`
	TakeProfit      string                 `json:"take_profit"`
	CumExecValue    string                 `json:"cum_exec_value"`
	CumExecFee      string                 `json:"cum_exec_fee"`
	RejectReason    string                 `json:"reject_reason"`
	OrderLinkID     string                 `json:"order_link_id"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	OrderID         string                 `json:"order_id"`
	Qty             float64                `json:"qty"`
	CumExecQty      int                    `json:"cum_exec_qty"`
	LeavesQty       int                    `json:"leaves_qty"`
	UserID          int                    `json:"user_id"`
	PositionIdx     int                    `json:"position_idx"`
}

// QueryStopOrderParam :
type QueryStopOrderParam struct {
	StopOrderID *string      `url:"stop_order_id,omitempty"`
	OrderLinkID *string      `url:"order_link_id,omitempty"`
	Symbol      SymbolFuture `url:"symbol"`
}

// QueryStopOrder :
func (s *FutureInversePerpetualService) QueryStopOrder(ctx context.Context, param QueryStopOrderParam) (*QueryStopOrderResponse, error) {
	var res QueryStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/v2/private/stop-order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// SaveLeverageResponse :
type SaveLeverageResponse struct {
	CommonResponse `json:",inline"`
	Result         float64 `json:"result"`
}

// SaveLeverageParam :
type SaveLeverageParam struct {
	Symbol   SymbolFuture `json:"symbol"`
	Leverage float64      `json:"leverage"`
}

// SaveLeverage :
func (s *FutureInversePerpetualService) SaveLeverage(ctx context.Context, param SaveLeverageParam) (*SaveLeverageResponse, error) {
	var res SaveLeverageResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/v2/private/position/leverage/save", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
