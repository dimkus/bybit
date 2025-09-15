package bybit

import (
	"bytes"
	"context"
	"errors"
	"github.com/mailru/easyjson"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/dimkus/websocket"
)

// V5WebsocketPublicServiceI :
type V5WebsocketPublicServiceI interface {
	Start(ErrHandler) error
	Run() error
	Ping() error
	Close() error

	// SubscribeOrderBook subscribes to the order book stream.
	// The response object passed to the callback is acquired from a sync.Pool.
	// Do not retain a reference to the object after the callback returns, as it will be recycled.
	SubscribeOrderBook(
		V5WebsocketPublicOrderBookParamKey,
		func(*V5WebsocketPublicOrderBookResponse) error,
	) (func() error, error)

	// SubscribeKline subscribes to the kline stream.
	// The response object passed to the callback is acquired from a sync.Pool.
	// Do not retain a reference to the object after the callback returns, as it will be recycled.
	SubscribeKline(
		V5WebsocketPublicKlineParamKey,
		func(*V5WebsocketPublicKlineResponse) error,
	) (func() error, error)

	// SubscribeTicker subscribes to the ticker stream.
	// The response object passed to the callback is acquired from a sync.Pool.
	// Do not retain a reference to the object after the callback returns, as it will be recycled.
	SubscribeTicker(
		V5WebsocketPublicTickerParamKey,
		func(*V5WebsocketPublicTickerResponse) error,
	) (func() error, error)

	// SubscribeTrade subscribes to the public trade stream.
	// The response object passed to the callback is acquired from a sync.Pool.
	// Do not retain a reference to the object after the callback returns, as it will be recycled.
	SubscribeTrade(
		V5WebsocketPublicTradeParamKey,
		func(*V5WebsocketPublicTradeResponse) error,
	) (func() error, error)

	// SubscribeLiquidation subscribes to the liquidation stream.
	// The response object passed to the callback is acquired from a sync.Pool.
	// Do not retain a reference to the object after the callback returns, as it will be recycled.
	SubscribeLiquidation(
		V5WebsocketPublicLiquidationParamKey,
		func(*V5WebsocketPublicLiquidationResponse) error,
	) (func() error, error)
}

// V5WebsocketPublicService :
type V5WebsocketPublicService struct {
	ctx                 context.Context
	client              *WebSocketClient
	connection          *websocket.Conn
	paramOrderBookMap   map[V5WebsocketPublicOrderBookParamKey]func(*V5WebsocketPublicOrderBookResponse) error
	paramKlineMap       map[V5WebsocketPublicKlineParamKey]func(*V5WebsocketPublicKlineResponse) error
	paramTickerMap      map[V5WebsocketPublicTickerParamKey]func(*V5WebsocketPublicTickerResponse) error
	paramTradeMap       map[V5WebsocketPublicTradeParamKey]func(*V5WebsocketPublicTradeResponse) error
	paramLiquidationMap map[V5WebsocketPublicLiquidationParamKey]func(*V5WebsocketPublicLiquidationResponse) error
	category            CategoryV5
	mu                  sync.Mutex
}

const (
	// V5WebsocketPublicPath :
	V5WebsocketPublicPath = "/v5/public"
)

// V5WebsocketPublicPathFor :
func V5WebsocketPublicPathFor(category CategoryV5) string {
	return V5WebsocketPublicPath + "/" + string(category)
}

// V5WebsocketPublicTopic :
type V5WebsocketPublicTopic string

const (
	// V5WebsocketPublicTopicOrderBook :
	V5WebsocketPublicTopicOrderBook = V5WebsocketPublicTopic("orderbook")

	// V5WebsocketPublicTopicKline :
	V5WebsocketPublicTopicKline = V5WebsocketPublicTopic("kline")

	// V5WebsocketPublicTopicTicker :
	V5WebsocketPublicTopicTicker = V5WebsocketPublicTopic("tickers")

	// V5WebsocketPublicTopicTrade :
	V5WebsocketPublicTopicTrade = V5WebsocketPublicTopic("publicTrade")

	// V5WebsocketPublicTopicLiquidation :
	V5WebsocketPublicTopicLiquidation = V5WebsocketPublicTopic("allLiquidation")
)

func (t V5WebsocketPublicTopic) String() string {
	return string(t)
}

// judgeTopic determines the topic of a websocket message without full JSON unmarshaling for performance.
func (v *V5WebsocketPublicService) judgeTopic(respBody []byte) (V5WebsocketPublicTopic, error) {
	// Search for the "topic" key in the byte slice. This is much faster than full JSON parsing.
	topicKey := []byte(`"topic":"`)
	idx := bytes.Index(respBody, topicKey)
	if idx == -1 {
		// "topic" key not found, this might not be a data message.
		return "", nil
	}

	// Move the starting position past the key to get the beginning of the value.
	start := idx + len(topicKey)

	// Find the end of the topic value (the next quote).
	end := bytes.IndexByte(respBody[start:], '"')
	if end == -1 {
		// Invalid format, no closing quote found.
		return "", errors.New("invalid topic format: missing closing quote")
	}

	// Extract the byte slice corresponding to the topic value without allocating a new string.
	topicSlice := respBody[start : start+end]

	// Compare the extracted byte slice with our constants.
	// bytes.Contains is faster than strings.Contains as it avoids string conversion.
	if bytes.Contains(topicSlice, []byte(V5WebsocketPublicTopicOrderBook)) {
		return V5WebsocketPublicTopicOrderBook, nil
	}
	if bytes.Contains(topicSlice, []byte(V5WebsocketPublicTopicKline)) {
		return V5WebsocketPublicTopicKline, nil
	}
	if bytes.Contains(topicSlice, []byte(V5WebsocketPublicTopicTicker)) {
		return V5WebsocketPublicTopicTicker, nil
	}
	if bytes.Contains(topicSlice, []byte(V5WebsocketPublicTopicTrade)) {
		return V5WebsocketPublicTopicTrade, nil
	}
	if bytes.Contains(topicSlice, []byte(V5WebsocketPublicTopicLiquidation)) {
		return V5WebsocketPublicTopicLiquidation, nil
	}

	return "", nil
}

// UnmarshalJSON :
func (v *V5WebsocketPublicTickerData) UnmarshalJSON(data []byte) error {
	switch v.category {
	case CategoryV5Linear, CategoryV5Inverse:
		return easyjson.Unmarshal(data, v.LinearInverse)
	case CategoryV5Option:
		return easyjson.Unmarshal(data, v.Option)
	case CategoryV5Spot:
		return easyjson.Unmarshal(data, v.Spot)
	}
	return errors.New("unsupported format")
}

// Start :
func (v *V5WebsocketPublicService) Start(errHandler ErrHandler) error {
	err := v.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
	if err != nil {
		return err
	}

	v.connection.SetPongHandler(func(string) error {
		return v.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
	})

	err = v.Run()
	if err != nil {
		return err
	}

	ctxToRun, ctxToRunCancel := signal.NotifyContext(v.ctx, os.Interrupt)

	go func(ctxToRun context.Context, s *V5WebsocketPublicService) {
		pingTicker := time.NewTicker(20 * time.Second)
		defer pingTicker.Stop()

		defer s.connection.Close()

		for {
			select {
			case <-ctxToRun.Done():
				s.client.debugf("caught websocket private service interrupt signal")
				closeErr := s.Close()
				if closeErr != nil {
					if errHandler == nil {
						return
					}
					errHandler(IsErrWebsocketClosed(closeErr), closeErr)
					return
				}
				return
			case <-pingTicker.C:
				pingErr := s.Ping()
				if pingErr != nil {
					ctxToRunCancel()

					if errHandler == nil {
						return
					}
					errHandler(IsErrWebsocketClosed(pingErr), pingErr)
					return
				}
			}
		}
	}(ctxToRun, v)

	go func(ctxToRun context.Context, s *V5WebsocketPublicService) {
		for {
			select {
			case <-ctxToRun.Done():
				s.client.debugf("caught websocket private service interrupt signal")
				closeErr := s.Close()
				if closeErr != nil {
					if errHandler == nil {
						return
					}
					errHandler(IsErrWebsocketClosed(closeErr), closeErr)
					return
				}
				return
			default:
				runErr := s.Run()
				if runErr != nil {
					ctxToRunCancel()
					if errHandler == nil {
						return
					}
					errHandler(IsErrWebsocketClosed(runErr), runErr)
					return
				}
			}
		}
	}(ctxToRun, v)

	return nil
}

var v5WebsocketPublicOrderBookResponseSyncPool = sync.Pool{
	New: func() interface{} {
		return &V5WebsocketPublicOrderBookResponse{
			Data: &V5WebsocketPublicOrderBookData{
				Bids: make([]PriceSizeTuple, 0, 10),
				Asks: make([]PriceSizeTuple, 0, 10),
			},
		}
	},
}

var v5WebsocketPublicKlineResponseSyncPool = sync.Pool{
	New: func() interface{} {
		return &V5WebsocketPublicKlineResponse{}
	},
}

var v5WebsocketPublicTickerResponseSyncPool = sync.Pool{
	New: func() interface{} {
		return &V5WebsocketPublicTickerResponse{
			Data: &V5WebsocketPublicTickerData{
				LinearInverse: &V5WebsocketPublicTickerLinearInverseResult{},
				Option:        &V5WebsocketPublicTickerOptionResult{},
				Spot:          &V5WebsocketPublicTickerSpotResult{},
			},
		}
	},
}

var v5WebsocketPublicTradeResponseSyncPool = sync.Pool{
	New: func() interface{} {
		return &V5WebsocketPublicTradeResponse{
			Data: make([]V5WebsocketPublicTradeData, 0, 10),
		}
	},
}

var v5WebsocketPublicLiquidationResponseSyncPool = sync.Pool{
	New: func() interface{} {
		return &V5WebsocketPublicLiquidationResponse{
			Data: make([]V5WebsocketPublicLiquidationData, 0, 10),
		}
	},
}

// Run
func (v *V5WebsocketPublicService) Run() error {
	var (
		err        error
		errProcess error
	)

	err = v.connection.ReadMessageWithPool(func(messageType int, message []byte) {
		topic, errJudgeTopic := v.judgeTopic(message)
		if errJudgeTopic != nil {
			errProcess = errJudgeTopic
			return
		}

		switch topic {
		case V5WebsocketPublicTopicOrderBook:
			resp := v5WebsocketPublicOrderBookResponseSyncPool.Get().(*V5WebsocketPublicOrderBookResponse)
			resp.clear()
			defer v5WebsocketPublicOrderBookResponseSyncPool.Put(resp)

			errUnmarshal := easyjson.Unmarshal(message, resp)
			if errUnmarshal != nil {
				errProcess = errUnmarshal
				return
			}

			f, errRetrieveOrderBookFunc := v.retrieveOrderBookFunc(resp.Key())
			if errRetrieveOrderBookFunc != nil {
				errProcess = errRetrieveOrderBookFunc
				return
			}
			if errF := f(resp); errF != nil {
				errProcess = errF
				return
			}
		case V5WebsocketPublicTopicKline:
			resp := v5WebsocketPublicKlineResponseSyncPool.Get().(*V5WebsocketPublicKlineResponse)
			resp.clear()
			defer v5WebsocketPublicKlineResponseSyncPool.Put(resp)

			errUnmarshal := easyjson.Unmarshal(message, resp)
			if errUnmarshal != nil {
				errProcess = errUnmarshal
				return
			}

			f, errRetrieveKlineFunc := v.retrieveKlineFunc(resp.Key())
			if errRetrieveKlineFunc != nil {
				errProcess = errRetrieveKlineFunc
				return
			}

			if errF := f(resp); errF != nil {
				errProcess = errF
				return
			}
		case V5WebsocketPublicTopicTicker:
			resp := v5WebsocketPublicTickerResponseSyncPool.Get().(*V5WebsocketPublicTickerResponse)
			resp.clear()
			defer v5WebsocketPublicTickerResponseSyncPool.Put(resp)

			resp.Data.category = v.category

			errUnmarshal := easyjson.Unmarshal(message, resp)
			if errUnmarshal != nil {
				errProcess = errUnmarshal
				return
			}

			f, errRetrieveTickerFunc := v.retrieveTickerFunc(resp.Key())
			if errRetrieveTickerFunc != nil {
				errProcess = errRetrieveTickerFunc
				return
			}

			if errF := f(resp); errF != nil {
				errProcess = errF
				return
			}
		case V5WebsocketPublicTopicTrade:
			resp := v5WebsocketPublicTradeResponseSyncPool.Get().(*V5WebsocketPublicTradeResponse)
			resp.clear()
			defer v5WebsocketPublicTradeResponseSyncPool.Put(resp)

			errUnmarshal := easyjson.Unmarshal(message, resp)
			if errUnmarshal != nil {
				errProcess = errUnmarshal
				return
			}

			f, errRetrieveTradeFunc := v.retrieveTradeFunc(resp.Key())
			if errRetrieveTradeFunc != nil {
				errProcess = errRetrieveTradeFunc
				return
			}

			if errF := f(resp); errF != nil {
				errProcess = errF
				return
			}
		case V5WebsocketPublicTopicLiquidation:
			resp := v5WebsocketPublicLiquidationResponseSyncPool.Get().(*V5WebsocketPublicLiquidationResponse)
			resp.clear()
			defer v5WebsocketPublicLiquidationResponseSyncPool.Put(resp)

			errUnmarshal := easyjson.Unmarshal(message, resp)
			if errUnmarshal != nil {
				errProcess = errUnmarshal
				return
			}

			f, errRetrieveLiquidationFunc := v.retrieveLiquidationFunc(resp.Key())
			if errRetrieveLiquidationFunc != nil {
				errProcess = errRetrieveLiquidationFunc
				return
			}

			if errF := f(resp); errF != nil {
				errProcess = errF
				return
			}
		}
	})

	if err != nil {
		return err
	}

	if errProcess != nil {
		return errProcess
	}

	return nil
}

// Ping :
func (v *V5WebsocketPublicService) Ping() error {
	// NOTE: It appears that two messages need to be sent.
	// REF: https://github.com/dimkus/bybit/pull/127#issuecomment-1537479346
	if err := v.writeMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	if err := v.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)); err != nil {
		return err
	}
	return nil
}

// Close :
func (v *V5WebsocketPublicService) Close() error {
	if err := v.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil && !errors.Is(err, websocket.ErrCloseSent) {
		return err
	}
	return nil
}

func (v *V5WebsocketPublicService) writeMessage(messageType int, body []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if err := v.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
