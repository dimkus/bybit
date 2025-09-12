package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FutureInverseFutureServiceI :
type FutureInverseFutureServiceI interface {
	// Market Data Endpoints
	OrderBook(context.Context, SymbolFuture) (*OrderBookResponse, error)
	ListKline(context.Context, ListKlineParam) (*ListKlineResponse, error)
	Tickers(context.Context, SymbolFuture) (*TickersResponse, error)
	TradingRecords(context.Context, TradingRecordsParam) (*TradingRecordsResponse, error)
	Symbols(context.Context) (*SymbolsResponse, error)
	MarkPriceKline(context.Context, MarkPriceKlineParam) (*MarkPriceKlineResponse, error)
	IndexPriceKline(context.Context, IndexPriceKlineParam) (*IndexPriceKlineResponse, error)
	OpenInterest(context.Context, OpenInterestParam) (*OpenInterestResponse, error)
	BigDeal(context.Context, BigDealParam) (*BigDealResponse, error)
	AccountRatio(context.Context, AccountRatioParam) (*AccountRatioResponse, error)

	// Account Data Endpoints
	CreateFuturesOrder(context.Context, CreateFuturesOrderParam) (*CreateFuturesOrderResponse, error)
	ListFuturesOrder(context.Context, ListFuturesOrderParam) (*ListFuturesOrderResponse, error)
	CancelFuturesOrder(context.Context, CancelFuturesOrderParam) (*CancelFuturesOrderResponse, error)
	CancelAllFuturesOrder(context.Context, CancelAllFuturesOrderParam) (*CancelAllFuturesOrderResponse, error)
	QueryFuturesOrder(context.Context, QueryFuturesOrderParam) (*QueryFuturesOrderResponse, error)
	CreateFuturesStopOrder(context.Context, CreateFuturesStopOrderParam) (*CreateFuturesStopOrderResponse, error)
	ListFuturesStopOrder(context.Context, ListFuturesStopOrderParam) (*ListFuturesStopOrderResponse, error)
	CancelFuturesStopOrder(context.Context, CancelFuturesStopOrderParam) (*CancelFuturesStopOrderResponse, error)
	CancelAllFuturesStopOrder(context.Context, CancelAllFuturesStopOrderParam) (*CancelAllFuturesStopOrderResponse, error)
	QueryFuturesStopOrder(context.Context, QueryFuturesStopOrderParam) (*QueryFuturesStopOrderResponse, error)
	ListFuturesPositions(context.Context, SymbolFuture) (*ListFuturesPositionsResponse, error)
	FuturesTradingStop(context.Context, FuturesTradingStopParam) (*FuturesTradingStopResponse, error)
	FuturesSaveLeverage(context.Context, FuturesSaveLeverageParam) (*FuturesSaveLeverageResponse, error)
	APIKeyInfo(context.Context) (*APIKeyInfoResponse, error)

	// Wallet Data Endpoints
	Balance(context.Context, Coin) (*BalanceResponse, error)
}

// FutureInverseFutureService :
type FutureInverseFutureService struct {
	client *Client

	*FutureCommonService
}

// CreateFuturesOrderResponse :
type CreateFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CreateFuturesOrderResult `json:"result"`
}

// CreateFuturesOrderResult :
type CreateFuturesOrderResult struct {
	TpTriggerBy   TriggerByFuture `json:"tp_trigger_by"`
	CreatedAt     string          `json:"created_at"`
	Symbol        SymbolFuture    `json:"symbol"`
	Side          Side            `json:"side"`
	OrderType     OrderType       `json:"order_type"`
	SlTriggerBy   TriggerByFuture `json:"sl_trigger_by"`
	OrderStatus   OrderStatus     `json:"order_status"`
	RejectReason  string          `json:"reject_reason"`
	OrderID       string          `json:"order_id"`
	StopLoss      string          `json:"stop_loss"`
	TakeProfit    string          `json:"take_profit"`
	UpdatedAt     string          `json:"updated_at"`
	TimeInForce   TimeInForce     `json:"time_in_force"`
	OrderLinkID   string          `json:"order_link_id"`
	Price         float64         `json:"price"`
	CumExecFee    float64         `json:"cum_exec_fee"`
	CumExecValue  float64         `json:"cum_exec_value"`
	CumExecQty    float64         `json:"cum_exec_qty"`
	LeavesQty     float64         `json:"leaves_qty"`
	LastExecPrice float64         `json:"last_exec_price"`
	LastExecTime  float64         `json:"last_exec_time"`
	Qty           float64         `json:"qty"`
	UserID        int             `json:"user_id"`
}

// CreateFuturesOrderParam :
type CreateFuturesOrderParam struct {
	OrderLinkID    *string          `json:"order_link_id,omitempty"`
	TakeProfit     *float64         `json:"take_profit,omitempty"`
	SlTriggerBy    *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	ReduceOnly     *bool            `json:"reduce_only,omitempty"`
	StopLoss       *float64         `json:"stop_loss,omitempty"`
	Price          *float64         `json:"price,omitempty"`
	CloseOnTrigger *bool            `json:"close_on_trigger,omitempty"`
	TpTriggerBy    *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	PositionIdx    *int             `json:"position_idx,omitempty"`
	Side           Side             `json:"side"`
	Symbol         SymbolFuture     `json:"symbol"`
	TimeInForce    TimeInForce      `json:"time_in_force"`
	OrderType      OrderType        `json:"order_type"`
	Qty            int              `json:"qty"`
}

// CreateFuturesOrder :
func (s *FutureInverseFutureService) CreateFuturesOrder(ctx context.Context, param CreateFuturesOrderParam) (*CreateFuturesOrderResponse, error) {
	var res CreateFuturesOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateFuturesOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListFuturesOrderResponse :
type ListFuturesOrderResponse struct {
	Result         ListFuturesOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListFuturesOrderResult :
type ListFuturesOrderResult struct {
	ListFuturesOrders []ListFuturesOrder `json:"data"`
}

// ListFuturesOrder :
type ListFuturesOrder struct {
	OrderStatus  OrderStatus     `json:"order_status"`
	Qty          string          `json:"qty"`
	Symbol       SymbolFuture    `json:"symbol"`
	Side         Side            `json:"side"`
	OrderType    OrderType       `json:"order_type"`
	Price        string          `json:"price"`
	UpdatedAt    string          `json:"updated_at"`
	TimeInForce  TimeInForce     `json:"time_in_force"`
	OrderLinkID  string          `json:"order_link_id"`
	OrderID      string          `json:"order_id"`
	Cursor       string          `json:"cursor"`
	CreatedAt    string          `json:"created_at"`
	CumExecValue string          `json:"cum_exec_value"`
	LeavesQty    string          `json:"leaves_qty"`
	LeavesValue  string          `json:"leaves_value"`
	CumExecQty   string          `json:"cum_exec_qty"`
	SlTriggerBy  TriggerByFuture `json:"sl_trigger_by"`
	CumExecFee   string          `json:"cum_exec_fee"`
	RejectReason string          `json:"reject_reason"`
	TakeProfit   string          `json:"take_profit"`
	StopLoss     string          `json:"stop_loss"`
	TpTriggerBy  TriggerByFuture `json:"tp_trigger_by"`
	UserID       int             `json:"user_id"`
	PositionIdx  int             `json:"position_idx"`
}

// ListFuturesOrderParam :
type ListFuturesOrderParam struct {
	OrderStatus *OrderStatus `url:"order_status,omitempty"`
	Direction   *Direction   `url:"direction,omitempty"`
	Limit       *int         `url:"limit,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
	Symbol      SymbolFuture `url:"symbol"`
}

// ListFuturesOrder :
func (s *FutureInverseFutureService) ListFuturesOrder(ctx context.Context, param ListFuturesOrderParam) (*ListFuturesOrderResponse, error) {
	var res ListFuturesOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/futures/private/order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelFuturesOrderResponse :
type CancelFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         CancelFuturesOrderResult `json:"result"`
}

// CancelFuturesOrderResult :
type CancelFuturesOrderResult struct {
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

// CancelFuturesOrderParam :
type CancelFuturesOrderParam struct {
	OrderID     *string      `json:"order_id,omitempty"`
	OrderLinkID *string      `json:"order_link_id,omitempty"`
	Symbol      SymbolFuture `json:"symbol"`
}

// CancelFuturesOrder :
func (s *FutureInverseFutureService) CancelFuturesOrder(ctx context.Context, param CancelFuturesOrderParam) (*CancelFuturesOrderResponse, error) {
	var res CancelFuturesOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelFuturesOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllFuturesOrderResponse :
type CancelAllFuturesOrderResponse struct {
	Result         []CancelAllFuturesOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CancelAllFuturesOrderResult :
type CancelAllFuturesOrderResult struct {
	CreateType  string       `json:"create_type"`
	CancelType  string       `json:"cancel_type"`
	Symbol      SymbolFuture `json:"symbol"`
	Side        Side         `json:"side"`
	OrderType   OrderType    `json:"order_type"`
	Price       string       `json:"price"`
	CrossStatus string       `json:"cross_status"`
	TimeInForce TimeInForce  `json:"time_in_force"`
	UpdatedAt   string       `json:"updated_at"`
	OrderStatus OrderStatus  `json:"order_status"`
	ClOrdID     string       `json:"clOrdID"`
	CreatedAt   string       `json:"created_at"`
	LeavesValue string       `json:"leaves_value"`
	LeavesQty   float64      `json:"leaves_qty"`
	UserID      int          `json:"user_id"`
	Qty         float64      `json:"qty"`
	CrossSeq    int          `json:"cross_seq"`
}

// CancelAllFuturesOrderParam :
type CancelAllFuturesOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllFuturesOrder :
func (s *FutureInverseFutureService) CancelAllFuturesOrder(ctx context.Context, param CancelAllFuturesOrderParam) (*CancelAllFuturesOrderResponse, error) {
	var res CancelAllFuturesOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllFuturesOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryFuturesOrderResponse :
type QueryFuturesOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         QueryFuturesOrderResult `json:"result"`
}

// QueryFuturesOrderResult :
type QueryFuturesOrderResult struct {
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

// QueryFuturesOrderParam :
type QueryFuturesOrderParam struct {
	OrderID     *string      `url:"order_id,omitempty"`
	OrderLinkID *string      `url:"order_link_id,omitempty"`
	Symbol      SymbolFuture `url:"symbol"`
}

// QueryFuturesOrder :
func (s *FutureInverseFutureService) QueryFuturesOrder(ctx context.Context, param QueryFuturesOrderParam) (*QueryFuturesOrderResponse, error) {
	var res QueryFuturesOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/futures/private/order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateFuturesStopOrderResponse :
type CreateFuturesStopOrderResponse struct {
	Result         CreateFuturesStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CreateFuturesStopOrderResult :
type CreateFuturesStopOrderResult struct {
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

// CreateFuturesStopOrderParam :
type CreateFuturesStopOrderParam struct {
	TakeProfit     *float64         `json:"take_profit,omitempty"`
	CloseOnTrigger *bool            `json:"close_on_trigger,omitempty"`
	SlTriggerBy    *TriggerByFuture `json:"sl_trigger_by,omitempty"`
	TpTriggerBy    *TriggerByFuture `json:"tp_trigger_by,omitempty"`
	StopLoss       *float64         `json:"stop_loss,omitempty"`
	OrderLinkID    *string          `json:"order_link_id,omitempty"`
	PositionIdx    *int             `json:"position_idx,omitempty"`
	TriggerBy      *TriggerByFuture `json:"trigger_by,omitempty"`
	Price          *float64         `json:"price,omitempty"`
	Symbol         SymbolFuture     `json:"symbol"`
	TimeInForce    TimeInForce      `json:"time_in_force"`
	Side           Side             `json:"side"`
	OrderType      OrderType        `json:"order_type"`
	StopPx         float64          `json:"stop_px"`
	BasePrice      float64          `json:"base_price"`
	Qty            float64          `json:"qty"`
}

// CreateFuturesStopOrder :
func (s *FutureInverseFutureService) CreateFuturesStopOrder(ctx context.Context, param CreateFuturesStopOrderParam) (*CreateFuturesStopOrderResponse, error) {
	var res CreateFuturesStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateFuturesStopOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/stop-order/create", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListFuturesStopOrderResponse :
type ListFuturesStopOrderResponse struct {
	Result         ListFuturesStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListFuturesStopOrderResult :
type ListFuturesStopOrderResult struct {
	ListFuturesStopOrders []ListFuturesStopOrder `json:"data"`
}

// ListFuturesStopOrder :
type ListFuturesStopOrder struct {
	BasePrice       string              `json:"base_price"`
	OrderLinkID     string              `json:"order_link_id"`
	StopOrderStatus OrderStatus         `json:"stop_order_status"`
	Symbol          SymbolFuture        `json:"symbol"`
	Side            Side                `json:"side"`
	OrderType       OrderType           `json:"order_type"`
	StopOrderType   StopOrderTypeFuture `json:"stop_order_type"`
	Price           string              `json:"price"`
	Qty             string              `json:"qty"`
	CreatedAt       string              `json:"created_at"`
	Cursor          string              `json:"cursor"`
	SlTriggerBy     TriggerByFuture     `json:"sl_trigger_by"`
	TimeInForce     TimeInForce         `json:"time_in_force"`
	UpdatedAt       string              `json:"updated_at"`
	StopPx          string              `json:"stop_px"`
	StopOrderID     string              `json:"stop_order_id"`
	TriggerBy       TriggerByFuture     `json:"trigger_by"`
	TakeProfit      string              `json:"take_profit"`
	StopLoss        string              `json:"stop_loss"`
	TpTriggerBy     TriggerByFuture     `json:"tp_trigger_by"`
	UserID          int                 `json:"user_id"`
	PositionIdx     int                 `json:"position_idx"`
}

// ListFuturesStopOrderParam :
type ListFuturesStopOrderParam struct {
	StopOrderStatus *OrderStatus `url:"stop_order_status,omitempty"`
	Direction       *Direction   `url:"direction,omitempty"`
	Limit           *int         `url:"limit,omitempty"`
	Cursor          *string      `url:"cursor,omitempty"`
	Symbol          SymbolFuture `url:"symbol"`
}

// ListFuturesStopOrder :
func (s *FutureInverseFutureService) ListFuturesStopOrder(ctx context.Context, param ListFuturesStopOrderParam) (*ListFuturesStopOrderResponse, error) {
	var res ListFuturesStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/futures/private/stop-order/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelFuturesStopOrderResponse :
type CancelFuturesStopOrderResponse struct {
	Result         CancelFuturesStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CancelFuturesStopOrderResult :
type CancelFuturesStopOrderResult struct {
	StopOrderID string `json:"stop_order_id"`
}

// CancelFuturesStopOrderParam :
type CancelFuturesStopOrderParam struct {
	StopOrderID *string      `json:"stop_order_id,omitempty"`
	OrderLinkID *string      `json:"order_link_id,omitempty"`
	Symbol      SymbolFuture `json:"symbol"`
}

// CancelFuturesStopOrder :
func (s *FutureInverseFutureService) CancelFuturesStopOrder(ctx context.Context, param CancelFuturesStopOrderParam) (*CancelFuturesStopOrderResponse, error) {
	var res CancelFuturesStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelFuturesStopOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/stop-order/cancel", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// CancelAllFuturesStopOrderResponse :
type CancelAllFuturesStopOrderResponse struct {
	Result         []CancelAllFuturesStopOrderResult `json:"result"`
	CommonResponse `json:",inline"`
}

// CancelAllFuturesStopOrderResult :
type CancelAllFuturesStopOrderResult struct {
	OrderStatus       OrderStatus         `json:"order_status"`
	CancelType        string              `json:"cancel_type"`
	Symbol            SymbolFuture        `json:"symbol"`
	Side              Side                `json:"side"`
	OrderType         OrderType           `json:"order_type"`
	Price             string              `json:"price"`
	ExpectedDirection string              `json:"expected_direction"`
	TimeInForce       TimeInForce         `json:"time_in_force"`
	CreateType        string              `json:"create_type"`
	StopOrderType     StopOrderTypeFuture `json:"stop_order_type"`
	BasePrice         string              `json:"base_price"`
	LeavesValue       string              `json:"leaves_value"`
	ClOrdID           string              `json:"clOrdID"`
	CreatedAt         string              `json:"created_at"`
	UpdatedAt         string              `json:"updated_at"`
	CrossStatus       string              `json:"cross_status"`
	TriggerBy         TriggerByFuture     `json:"trigger_by"`
	LeavesQty         float64             `json:"leaves_qty"`
	CrossSeq          int                 `json:"cross_seq"`
	UserID            int                 `json:"user_id"`
	Qty               float64             `json:"qty"`
}

// CancelAllFuturesStopOrderParam :
type CancelAllFuturesStopOrderParam struct {
	Symbol SymbolFuture `json:"symbol"`
}

// CancelAllFuturesStopOrder :
func (s *FutureInverseFutureService) CancelAllFuturesStopOrder(ctx context.Context, param CancelAllFuturesStopOrderParam) (*CancelAllFuturesStopOrderResponse, error) {
	var res CancelAllFuturesStopOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CancelAllFuturesStopOrderParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/stop-order/cancelAll", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// QueryFuturesStopOrderResponse :
type QueryFuturesStopOrderResponse struct {
	CommonResponse `json:",inline"`
	Result         QueryFuturesStopOrderResult `json:"result"`
}

// QueryFuturesStopOrderResult :
type QueryFuturesStopOrderResult struct {
	ExtFields    map[string]interface{} `json:"ext_fields"`
	OrderStatus  OrderStatus            `json:"order_status"`
	StopLoss     string                 `json:"stop_loss"`
	Side         Side                   `json:"side"`
	OrderType    OrderType              `json:"order_type"`
	Price        string                 `json:"price"`
	TpTriggerBy  TriggerByFuture        `json:"tp_trigger_by"`
	StopPx       string                 `json:"stop_px"`
	BasePrice    string                 `json:"base_price"`
	TimeInForce  TimeInForce            `json:"time_in_force"`
	TriggerBy    TriggerByFuture        `json:"trigger_by"`
	Symbol       SymbolFuture           `json:"symbol"`
	LeavesValue  string                 `json:"leaves_value"`
	SlTriggerBy  TriggerByFuture        `json:"sl_trigger_by"`
	TakeProfit   string                 `json:"take_profit"`
	CumExecValue string                 `json:"cum_exec_value"`
	CumExecFee   string                 `json:"cum_exec_fee"`
	RejectReason string                 `json:"reject_reason"`
	OrderLinkID  string                 `json:"order_link_id"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	OrderID      string                 `json:"order_id"`
	Qty          float64                `json:"qty"`
	CumExecQty   int                    `json:"cum_exec_qty"`
	LeavesQty    int                    `json:"leaves_qty"`
	UserID       int                    `json:"user_id"`
	PositionIdx  int                    `json:"position_idx"`
}

// QueryFuturesStopOrderParam :
type QueryFuturesStopOrderParam struct {
	StopOrderID *string      `url:"stop_order_id,omitempty"`
	OrderLinkID *string      `url:"order_link_id,omitempty"`
	Symbol      SymbolFuture `url:"symbol"`
}

// QueryFuturesStopOrder :
func (s *FutureInverseFutureService) QueryFuturesStopOrder(ctx context.Context, param QueryFuturesStopOrderParam) (*QueryFuturesStopOrderResponse, error) {
	var res QueryFuturesStopOrderResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPrivately(ctx, "/futures/private/stop-order", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// ListFuturesPositionsResponse :
type ListFuturesPositionsResponse struct {
	Result         []ListFuturesPositionsResult `json:"result"`
	CommonResponse `json:",inline"`
}

// ListFuturesPositionsResult :
type ListFuturesPositionsResult struct {
	Data ListFuturesPositionsResultData `json:"data"`
}

// ListFuturesPositionsResultData :
type ListFuturesPositionsResultData struct {
	BustPrice           string       `json:"bust_price"`
	OccClosingFee       string       `json:"occ_closing_fee"`
	TpSlMode            TpSlMode     `json:"tp_sl_mode"`
	UpdatedAt           string       `json:"updated_at"`
	CreatedAt           string       `json:"created_at"`
	Symbol              SymbolFuture `json:"symbol"`
	Side                Side         `json:"side"`
	CumRealisedPnl      string       `json:"cum_realised_pnl"`
	PositionValue       string       `json:"position_value"`
	EntryPrice          string       `json:"entry_price"`
	RealisedPnl         string       `json:"realised_pnl"`
	WalletBalance       string       `json:"wallet_balance"`
	Leverage            string       `json:"leverage"`
	EffectiveLeverage   string       `json:"effective_leverage"`
	PositionMargin      string       `json:"position_margin"`
	LiqPrice            string       `json:"liq_price"`
	OrderMargin         string       `json:"order_margin"`
	OcCalcData          string       `json:"oc_calc_data"`
	OccFundingFee       string       `json:"occ_funding_fee"`
	TakeProfit          string       `json:"take_profit"`
	StopLoss            string       `json:"stop_loss"`
	TrailingStop        string       `json:"trailing_stop"`
	PositionStatus      string       `json:"position_status"`
	DeleverageIndicator int          `json:"deleverage_indicator"`
	ID                  int          `json:"id"`
	PositionIdx         int          `json:"position_idx"`
	AutoAddMargin       float64      `json:"auto_add_margin"`
	UnrealisedPnl       float64      `json:"unrealised_pnl"`
	Size                float64      `json:"size"`
	CrossSeq            float64      `json:"cross_seq"`
	PositionSeq         float64      `json:"position_seq"`
	RiskID              int          `json:"risk_id"`
	UserID              int          `json:"user_id"`
	Mode                int          `json:"mode"`
	IsIsolated          bool         `json:"is_isolated"`
}

// ListFuturesPositions :
func (s *FutureInverseFutureService) ListFuturesPositions(ctx context.Context, symbol SymbolFuture) (*ListFuturesPositionsResponse, error) {
	var res ListFuturesPositionsResponse

	query := url.Values{}
	query.Add("symbol", string(symbol))

	if err := s.client.getPrivately(ctx, "/futures/private/position/list", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FuturesTradingStopResponse :
type FuturesTradingStopResponse struct {
	CommonResponse `json:",inline"`
	Result         FuturesTradingStopResult `json:"result"`
}

// FuturesTradingStopResult :
type FuturesTradingStopResult struct {
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

// FuturesTradingStopParam :
type FuturesTradingStopParam struct {
	PositionIdx       *int             `json:"position_idx,omitempty"`
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

// FuturesTradingStop :
func (s *FutureInverseFutureService) FuturesTradingStop(ctx context.Context, param FuturesTradingStopParam) (*FuturesTradingStopResponse, error) {
	var res FuturesTradingStopResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for FuturesTradingStopParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/position/trading-stop", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FuturesSaveLeverageResponse :
type FuturesSaveLeverageResponse struct {
	CommonResponse `json:",inline"`
	Result         float64 `json:"result"`
}

// FuturesSaveLeverageParam :
type FuturesSaveLeverageParam struct {
	Symbol       SymbolFuture `json:"symbol"`
	BuyLeverage  float64      `json:"buy_leverage"`
	SellLeverage float64      `json:"sell_leverage"`
}

// FuturesSaveLeverage :
func (s *FutureInverseFutureService) FuturesSaveLeverage(ctx context.Context, param FuturesSaveLeverageParam) (*FuturesSaveLeverageResponse, error) {
	var res FuturesSaveLeverageResponse

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for FuturesSaveLeverageParam: %w", err)
	}

	if err := s.client.postJSON(ctx, "/futures/private/position/leverage/save", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
