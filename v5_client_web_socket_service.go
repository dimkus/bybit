package bybit

import (
	"context"
	"github.com/gorilla/websocket"
)

// V5WebsocketServiceI :
type V5WebsocketServiceI interface {
	Public(ctx context.Context, category CategoryV5) (V5WebsocketPublicServiceI, error)
	Private(ctx context.Context) (V5WebsocketPrivateServiceI, error)
}

// V5WebsocketService :
type V5WebsocketService struct {
	client *WebSocketClient
}

// Public :
// Establishes a public websocket connection to Bybit V5.
//
// category is the category of the websocket connection.
//
// Returns a V5WebsocketPublicServiceI if successful.
func (s *V5WebsocketService) Public(ctx context.Context, category CategoryV5) (V5WebsocketPublicServiceI, error) {
	url := s.client.baseURL + V5WebsocketPublicPathFor(category)
	c, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPublicService{
		client:              s.client,
		connection:          c,
		category:            category,
		paramOrderBookMap:   make(map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error),
		paramKlineMap:       make(map[V5WebsocketPublicKlineParamKey]func(V5WebsocketPublicKlineResponse) error),
		paramTickerMap:      make(map[V5WebsocketPublicTickerParamKey]func(V5WebsocketPublicTickerResponse) error),
		paramTradeMap:       make(map[V5WebsocketPublicTradeParamKey]func(V5WebsocketPublicTradeResponse) error),
		paramLiquidationMap: make(map[V5WebsocketPublicLiquidationParamKey]func(V5WebsocketPublicLiquidationResponse) error),
	}, nil
}

// Private establishes a private websocket connection to Bybit V5.
//
// This method dials a websocket connection to the private endpoint using the
// base URL from the WebSocketClient and returns a V5WebsocketPrivateServiceI
// interface if successful.
//
// ctx is the context to control the lifecycle of the request.
//
// Returns a V5WebsocketPrivateServiceI if successful. If there is an error
// during the connection, it returns a non-nil error.
func (s *V5WebsocketService) Private(ctx context.Context) (V5WebsocketPrivateServiceI, error) {
	url := s.client.baseURL + V5WebsocketPrivatePath
	c, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return &V5WebsocketPrivateService{
		client:            s.client,
		ctx:               ctx,
		connection:        c,
		paramOrderMap:     make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error),
		paramPositionMap:  make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error),
		paramExecutionMap: make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateExecutionResponse) error),
		paramWalletMap:    make(map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error),
	}, nil
}

// V5 :
func (c *WebSocketClient) V5() *V5WebsocketService {
	return &V5WebsocketService{c}
}
