package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dimkus/websocket"
)

// SubscribeOrderBook :
func (v *V5WebsocketPublicService) SubscribeOrderBook(
	key V5WebsocketPublicOrderBookParamKey,
	f func(*V5WebsocketPublicOrderBookResponse) error,
) (func() error, error) {
	if err := v.addParamOrderBookFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{key.Topic()},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	if err := v.writeMessage(websocket.TextMessage, buf); err != nil {
		return nil, err
	}
	return func() error {
		param := struct {
			Op   string        `json:"op"`
			Args []interface{} `json:"args"`
		}{
			Op:   "unsubscribe",
			Args: []interface{}{key.Topic()},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := v.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		v.removeParamOrderBookFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicOrderBookParamKey :
type V5WebsocketPublicOrderBookParamKey struct {
	Symbol SymbolV5
	Depth  int
}

// Topic :
func (k *V5WebsocketPublicOrderBookParamKey) Topic() string {
	return fmt.Sprintf("%s.%d.%s", V5WebsocketPublicTopicOrderBook, k.Depth, k.Symbol)
}

// V5WebsocketPublicOrderBookResponse :
// easyjson:json
type V5WebsocketPublicOrderBookResponse struct {
	Data      *V5WebsocketPublicOrderBookData `json:"data"`
	Topic     string                          `json:"topic"`
	Type      string                          `json:"type"`
	TimeStamp int64                           `json:"ts"`
}

// clear :
func (r *V5WebsocketPublicOrderBookResponse) clear() {
	r.Topic = ""
	r.Type = ""
	r.TimeStamp = 0
	if r.Data != nil {
		r.Data.clear()
	}
}

// V5WebsocketPublicOrderBookData :
// easyjson:json
type V5WebsocketPublicOrderBookData struct {
	Symbol   SymbolV5           `json:"s"`
	Bids     PriceSizeTupleList `json:"b"`
	Asks     PriceSizeTupleList `json:"a"`
	UpdateID int                `json:"u"`
	Seq      int                `json:"seq"`
}

// PriceSizeTupleList
//
//easyjson:json
type PriceSizeTupleList []PriceSizeTuple

// PriceSizeTuple
//
//easyjson:json
type PriceSizeTuple [2]string

// GetPrice :
func (p *PriceSizeTuple) GetPrice() string {
	return p[0]
}

// GetSize :
func (p *PriceSizeTuple) GetSize() string {
	return p[1]
}

// clear :
func (p *PriceSizeTupleList) clear() {
	const maxCapacity = 40
	if cap(*p) > maxCapacity {
		*p = nil
		return
	}
	*p = (*p)[:0]
}

// clear :
func (r *V5WebsocketPublicOrderBookData) clear() {
	r.Symbol = ""
	r.UpdateID = 0
	r.Seq = 0
	if r.Bids != nil {
		r.Bids.clear()
	}

	if r.Asks != nil {
		r.Asks.clear()
	}
}

// Key :
func (r *V5WebsocketPublicOrderBookResponse) Key() V5WebsocketPublicOrderBookParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicOrderBook.String() || len(arr) != 3 {
		return V5WebsocketPublicOrderBookParamKey{}
	}
	depth, err := strconv.Atoi(arr[1])
	if err != nil {
		return V5WebsocketPublicOrderBookParamKey{}
	}
	symbol := SymbolV5(arr[2])
	return V5WebsocketPublicOrderBookParamKey{
		Depth:  depth,
		Symbol: symbol,
	}
}

// addParamOrderBookFunc :
func (v *V5WebsocketPublicService) addParamOrderBookFunc(key V5WebsocketPublicOrderBookParamKey, f func(*V5WebsocketPublicOrderBookResponse) error) error {
	if _, exist := v.paramOrderBookMap[key]; exist {
		return errors.New("already registered for this param")
	}
	v.paramOrderBookMap[key] = f
	return nil
}

// removeParamTradeFunc :
func (v *V5WebsocketPublicService) removeParamOrderBookFunc(key V5WebsocketPublicOrderBookParamKey) {
	delete(v.paramOrderBookMap, key)
}

// retrievePositionFunc :
func (v *V5WebsocketPublicService) retrieveOrderBookFunc(key V5WebsocketPublicOrderBookParamKey) (func(*V5WebsocketPublicOrderBookResponse) error, error) {
	f, exist := v.paramOrderBookMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
