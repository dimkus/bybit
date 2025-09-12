package bybit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

// V5PositionServiceI :
type V5PositionServiceI interface {
	GetPositionInfo(context.Context, V5GetPositionInfoParam) (*V5GetPositionInfoResponse, error)
	SetLeverage(context.Context, V5SetLeverageParam) (*V5SetLeverageResponse, error)
	SetTradingStop(context.Context, V5SetTradingStopParam) (*V5SetTradingStopResponse, error)
	SetTpSlMode(context.Context, V5SetTpSlModeParam) (*V5SetTpSlModeResponse, error)
	SwitchPositionMode(context.Context, V5SwitchPositionModeParam) (*V5SwitchPositionModeResponse, error)
	GetClosedPnL(context.Context, V5GetClosedPnLParam) (*V5GetClosedPnLResponse, error)
	SwitchPositionMarginMode(context.Context, V5SwitchPositionMarginModeParam) (*V5SwitchPositionMarginModeResponse, error)
	SetRiskLimit(context.Context, V5SetRiskLimitParam) (*V5SetRiskLimitResponse, error)
}

// V5PositionService :
type V5PositionService struct {
	client *Client
}

// V5GetPositionInfoParam :
type V5GetPositionInfoParam struct {
	Symbol     *SymbolV5  `url:"symbol,omitempty"`
	BaseCoin   *Coin      `url:"baseCoin,omitempty"`
	SettleCoin *Coin      `url:"settleCoin,omitempty"`
	Limit      *int       `url:"limit,omitempty"`
	Cursor     *string    `url:"cursor,omitempty"`
	Category   CategoryV5 `url:"category"`
}

// V5GetPositionInfoResponse :
type V5GetPositionInfoResponse struct {
	Result           V5GetPositionInfoResult `json:"result"`
	CommonV5Response `json:",inline"`
}

// V5GetPositionInfoResult :
type V5GetPositionInfoResult struct {
	Category       CategoryV5            `json:"category"`
	NextPageCursor string                `json:"nextPageCursor"`
	List           V5GetPositionInfoList `json:"list"`
}

// V5GetPositionInfoList :
type V5GetPositionInfoList []V5GetPositionInfoItem

// V5GetPositionInfoItem :
type V5GetPositionInfoItem struct {
	CumRealisedPnl         string           `json:"cumRealisedPnl"`
	PositionValue          string           `json:"positionValue"`
	AvgPrice               string           `json:"avgPrice"`
	LiqPrice               string           `json:"liqPrice"`
	RiskLimitValue         string           `json:"riskLimitValue"`
	TakeProfit             string           `json:"takeProfit"`
	MarkPrice              string           `json:"markPrice"`
	TpSlMode               TpSlMode         `json:"tpslMode"`
	StopLoss               string           `json:"stopLoss"`
	TrailingStop           string           `json:"trailingStop"`
	Leverage               string           `json:"leverage"`
	UnrealisedPnl          string           `json:"unrealisedPnl"`
	LeverageSysUpdatedTime string           `json:"leverageSysUpdatedTime"`
	PositionStatus         string           `json:"positionStatus"`
	Size                   string           `json:"size"`
	MmrSysUpdatedTime      string           `json:"mmrSysUpdatedTime"`
	Side                   Side             `json:"side"`
	PositionMM             string           `json:"positionMM"`
	PositionBalance        string           `json:"positionBalance"`
	CreatedTime            string           `json:"createdTime"`
	UpdatedTime            string           `json:"updatedTime"`
	Symbol                 SymbolV5         `json:"symbol"`
	BustPrice              string           `json:"bustPrice"`
	PositionIM             string           `json:"positionIM"`
	Seq                    int64            `json:"seq"`
	PositionIdx            int              `json:"positionIdx"`
	AdlRankIndicator       AdlRankIndicator `json:"adlRankIndicator"`
	RiskID                 int              `json:"riskId"`
	TradeMode              int              `json:"tradeMode"`
	AutoAddMargin          int              `json:"autoAddMargin"`
	IsReduceOnly           bool             `json:"isReduceOnly"`
}

// GetPositionInfo :
func (s *V5PositionService) GetPositionInfo(ctx context.Context, param V5GetPositionInfoParam) (*V5GetPositionInfoResponse, error) {
	var res V5GetPositionInfoResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately(ctx, "/v5/position/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5SetLeverageParam :
type V5SetLeverageParam struct {
	Category     CategoryV5 `json:"category"`
	Symbol       SymbolV5   `json:"symbol"`
	BuyLeverage  string     `json:"buyLeverage"`
	SellLeverage string     `json:"sellLeverage"`
}

// V5SetLeverageResponse :
type V5SetLeverageResponse struct {
	Result           interface{} `json:"result"`
	CommonV5Response `json:",inline"`
}

// SetLeverage :
func (s *V5PositionService) SetLeverage(ctx context.Context, param V5SetLeverageParam) (*V5SetLeverageResponse, error) {
	var res V5SetLeverageResponse

	if param.Category == "" || param.Symbol == "" || param.BuyLeverage == "" || param.SellLeverage == "" {
		return nil, fmt.Errorf("category, symbol, BuyLeverage and SellLeverage needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/position/set-leverage", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5SetTradingStopParam :
type V5SetTradingStopParam struct {
	ActivePrice  *string     `json:"activePrice,omitempty"`
	TpSize       *string     `json:"tpSize,omitempty"`
	SlOrderType  *OrderType  `json:"slOrderType,omitempty"`
	TakeProfit   *string     `json:"takeProfit,omitempty"`
	StopLoss     *string     `json:"stopLoss,omitempty"`
	TrailingStop *string     `json:"trailingStop,omitempty"`
	TpOrderType  *OrderType  `json:"tpOrderType,omitempty"`
	TpTriggerBy  *TriggerBy  `json:"tpTriggerBy,omitempty"`
	SlLimitPrice *string     `json:"slLimitPrice,omitempty"`
	SlTriggerBy  *TriggerBy  `json:"slTriggerBy,omitempty"`
	SlSize       *string     `json:"slSize,omitempty"`
	TpslMode     *TpSlMode   `json:"tpslMode,omitempty"`
	TpLimitPrice *string     `json:"tpLimitPrice,omitempty"`
	Category     CategoryV5  `json:"category"`
	Symbol       SymbolV5    `json:"symbol"`
	PositionIdx  PositionIdx `json:"positionIdx"`
}

func (p V5SetTradingStopParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	if p.TakeProfit == nil && p.StopLoss == nil && p.TpLimitPrice == nil && p.SlLimitPrice == nil {
		return fmt.Errorf("takeProfit or stopLoss or tpLimitPrice or slLimitPrice needed")
	}
	return nil
}

// V5SetTradingStopResponse :
type V5SetTradingStopResponse struct {
	Result           interface{} `json:"result"`
	CommonV5Response `json:",inline"`
}

// SetTradingStop :
func (s *V5PositionService) SetTradingStop(ctx context.Context, param V5SetTradingStopParam) (*V5SetTradingStopResponse, error) {
	var res V5SetTradingStopResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/position/trading-stop", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5SetTpSlModeParam :
type V5SetTpSlModeParam struct {
	Category CategoryV5 `json:"category"`
	Symbol   SymbolV5   `json:"symbol"`
	TpSlMode TpSlMode   `json:"tpSlMode"`
}

func (p V5SetTpSlModeParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	return nil
}

// V5SetTpSlModeResponse :
type V5SetTpSlModeResponse struct {
	Result           V5SetTpSlModeResult `json:"result"`
	CommonV5Response `json:",inline"`
}

// V5SetTpSlModeResult :
type V5SetTpSlModeResult struct {
	TpSlMode TpSlMode `json:"tpSlMode"`
}

// SetTpSlMode :
func (s *V5PositionService) SetTpSlMode(ctx context.Context, param V5SetTpSlModeParam) (*V5SetTpSlModeResponse, error) {
	var res V5SetTpSlModeResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/position/set-tpsl-mode", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5SwitchPositionModeParam :
type V5SwitchPositionModeParam struct {
	Symbol   *SymbolV5    `json:"symbol,omitempty"`
	Coin     *Coin        `json:"coin,omitempty"`
	Category CategoryV5   `json:"category"`
	Mode     PositionMode `json:"mode"`
}

func (p V5SwitchPositionModeParam) validate() error {
	if p.Symbol == nil && p.Coin == nil {
		return fmt.Errorf("symbol or coin is required")
	}
	return nil
}

// V5SwitchPositionModeResponse :
type V5SwitchPositionModeResponse struct {
	Result           interface{} `json:"result"`
	CommonV5Response `json:",inline"`
}

// SwitchPositionMode :
func (s *V5PositionService) SwitchPositionMode(ctx context.Context, param V5SwitchPositionModeParam) (*V5SwitchPositionModeResponse, error) {
	var res V5SwitchPositionModeResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/position/switch-mode", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5GetClosedPnLParam :
type V5GetClosedPnLParam struct {
	Symbol    *SymbolV5  `url:"symbol,omitempty"`
	StartTime *int64     `url:"startTime,omitempty"`
	EndTime   *int64     `url:"endTime,omitempty"`
	Limit     *int       `url:"limit,omitempty"`
	Cursor    *string    `url:"cursor,omitempty"`
	Category  CategoryV5 `url:"category"`
}

// V5GetClosedPnLResponse :
type V5GetClosedPnLResponse struct {
	Result           V5GetClosedPnLResult `json:"result"`
	CommonV5Response `json:",inline"`
}

// V5GetClosedPnLResult :
type V5GetClosedPnLResult struct {
	Category       CategoryV5         `json:"category"`
	NextPageCursor string             `json:"nextPageCursor"`
	List           V5GetClosedPnLList `json:"list"`
}

// V5GetClosedPnLList :
type V5GetClosedPnLList []V5GetClosedPnLItem

// V5GetClosedPnLItem :
type V5GetClosedPnLItem struct {
	Symbol        SymbolV5   `json:"symbol"`
	OrderID       string     `json:"orderId"`
	Side          Side       `json:"side"`
	Qty           string     `json:"qty"`
	OrderPrice    string     `json:"orderPrice"`
	OrderType     OrderType  `json:"orderType"`
	ExecType      ExecTypeV5 `json:"execType"`
	ClosedSize    string     `json:"closedSize"`
	CumEntryValue string     `json:"cumEntryValue"`
	AvgEntryPrice string     `json:"avgEntryPrice"`
	CumExitValue  string     `json:"cumExitValue"`
	AvgExitPrice  string     `json:"avgExitPrice"`
	ClosedPnl     string     `json:"closedPnl"`
	FillCount     string     `json:"fillCount"`
	Leverage      string     `json:"leverage"`
	CreatedTime   string     `json:"createdTime"`
	UpdatedTime   string     `json:"updatedTime"`
}

// GetClosedPnL :
func (s *V5PositionService) GetClosedPnL(ctx context.Context, param V5GetClosedPnLParam) (*V5GetClosedPnLResponse, error) {
	var res V5GetClosedPnLResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately(ctx, "/v5/position/closed-pnl", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5SwitchPositionMarginModeParam :
type V5SwitchPositionMarginModeParam struct {
	Category     CategoryV5         `json:"category"`
	Symbol       SymbolV5           `json:"symbol"`
	BuyLeverage  string             `json:"buyLeverage"`
	SellLeverage string             `json:"sellLeverage"`
	TradeMode    PositionMarginMode `json:"tradeMode"`
}

func (p V5SwitchPositionMarginModeParam) validate() error {
	if p.Category != CategoryV5Linear && p.Category != CategoryV5Inverse {
		return fmt.Errorf("only linear and inverse are supported for category")
	}
	if p.Symbol == "" || p.BuyLeverage == "" || p.SellLeverage == "" {
		return fmt.Errorf("symbol, BuyLeverage and SellLeverage needed")
	}
	return nil
}

// V5SwitchPositionMarginModeResponse :
type V5SwitchPositionMarginModeResponse struct {
	Result           interface{} `json:"result"`
	CommonV5Response `json:",inline"`
}

// SwitchPositionMarginMode :
func (s *V5PositionService) SwitchPositionMarginMode(ctx context.Context, param V5SwitchPositionMarginModeParam) (*V5SwitchPositionMarginModeResponse, error) {
	var res V5SwitchPositionMarginModeResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/position/switch-isolated", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

// V5SetRiskLimitParam :
type V5SetRiskLimitParam struct {
	PositionIdx *PositionIdx `json:"positionIdx,omitempty"`
	Category    CategoryV5   `json:"category"`
	Symbol      SymbolV5     `json:"symbol"`
	RiskID      int64        `json:"riskId"`
}

// V5SetRiskLimitResponse :
type V5SetRiskLimitResponse struct {
	Result           V5SetRiskLimitResult `json:"result"`
	CommonV5Response `json:",inline"`
}

// V5SetRiskLimitResult :
type V5SetRiskLimitResult struct {
	Category       CategoryV5 `json:"category"`
	RiskLimitValue string     `json:"riskLimitValue"`
	RiskID         int64      `json:"riskId"`
}

// SetRiskLimit :
func (s *V5PositionService) SetRiskLimit(ctx context.Context, param V5SetRiskLimitParam) (*V5SetRiskLimitResponse, error) {
	var res V5SetRiskLimitResponse

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/position/set-risk-limit", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}
