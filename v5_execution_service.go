package bybit

import (
	"context"
	"github.com/google/go-querystring/query"
)

// V5ExecutionServiceI :
type V5ExecutionServiceI interface {
	GetExecutionList(context.Context, V5GetExecutionParam) (*V5GetExecutionListResponse, error)
}

// V5ExecutionService :
type V5ExecutionService struct {
	client *Client
}

// V5GetExecutionParam :
type V5GetExecutionParam struct {
	Symbol      *SymbolV5   `url:"symbol,omitempty"`
	OrderID     *string     `url:"orderId,omitempty"`
	OrderLinkID *string     `url:"orderLinkId,omitempty"`
	BaseCoin    *Coin       `url:"baseCoin,omitempty"`
	StartTime   *int        `url:"startTime,omitempty"`
	EndTime     *int        `url:"endTime,omitempty"`
	ExecType    *ExecTypeV5 `url:"execType,omitempty"`
	Limit       *int        `url:"limit,omitempty"`
	Cursor      *string     `url:"cursor,omitempty"`
	Category    CategoryV5  `url:"category"`
}

// V5GetExecutionListResponse :
type V5GetExecutionListResponse struct {
	Result           V5GetExecutionListResult `json:"result"`
	CommonV5Response `json:",inline"`
}

// V5GetExecutionListResult :
type V5GetExecutionListResult struct {
	NextPageCursor string                   `json:"nextPageCursor"`
	Category       string                   `json:"category"`
	List           []V5GetExecutionListItem `json:"list"`
}

// V5GetExecutionListItem :
type V5GetExecutionListItem struct {
	ExecQty         string     `json:"execQty"`
	IndexPrice      string     `json:"indexPrice"`
	OrderLinkID     string     `json:"orderLinkId"`
	Side            Side       `json:"side"`
	OrderPrice      string     `json:"orderPrice"`
	OrderQty        string     `json:"orderQty"`
	LeavesQty       string     `json:"leavesQty"`
	OrderType       OrderType  `json:"orderType"`
	StopOrderType   string     `json:"stopOrderType"`
	ExecFee         string     `json:"execFee"`
	ExecID          string     `json:"execId"`
	ExecPrice       string     `json:"execPrice"`
	OrderID         string     `json:"orderId"`
	ExecValue       string     `json:"execValue"`
	Symbol          SymbolV5   `json:"symbol"`
	ExecTime        string     `json:"execTime"`
	ClosedSize      string     `json:"closedSize"`
	FeeRate         string     `json:"feeRate"`
	TradeIv         string     `json:"tradeIv"`
	MarkIv          string     `json:"markIv"`
	MarkPrice       string     `json:"markPrice"`
	ExecType        ExecTypeV5 `json:"execType"`
	UnderlyingPrice string     `json:"underlyingPrice"`
	BlockTradeID    string     `json:"blockTradeId"`
	IsMaker         bool       `json:"isMaker"`
}

func (s *V5ExecutionService) GetExecutionList(ctx context.Context, param V5GetExecutionParam) (*V5GetExecutionListResponse, error) {
	var res V5GetExecutionListResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5Privately(ctx, "/v5/execution/list", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
