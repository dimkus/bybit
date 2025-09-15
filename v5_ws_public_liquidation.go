package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dimkus/websocket"
)

// SubscribeLiquidation :
// IMPORTANT: This function is not thread-safe
func (v *V5WebsocketPublicService) SubscribeLiquidation(
	key V5WebsocketPublicLiquidationParamKey,
	f func(*V5WebsocketPublicLiquidationResponse) error,
) (func() error, error) {
	if err := v.addParamLiquidationFunc(key, f); err != nil {
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
		v.removeParamLiquidationFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicLiquidationParamKey :
type V5WebsocketPublicLiquidationParamKey struct {
	Symbol SymbolV5
}

// Topic :
func (k *V5WebsocketPublicLiquidationParamKey) Topic() string {
	return fmt.Sprintf("%s.%s", V5WebsocketPublicTopicLiquidation, k.Symbol)
}

// V5WebsocketPublicLiquidationResponse :
// easyjson:json
type V5WebsocketPublicLiquidationResponse struct {
	Topic     string                             `json:"topic"`
	Type      string                             `json:"type"`
	Data      []V5WebsocketPublicLiquidationData `json:"data"`
	TimeStamp int64                              `json:"ts"`
}

// clear :
func (r *V5WebsocketPublicLiquidationResponse) clear() {
	r.Topic = ""
	r.Type = ""
	r.TimeStamp = 0
	if cap(r.Data) > 40 {
		r.Data = nil
	} else {
		for i := range r.Data {
			r.Data[i] = V5WebsocketPublicLiquidationData{}
		}
		r.Data = r.Data[:0]
	}
}

// V5WebsocketPublicLiquidationData :
// easyjson:json
type V5WebsocketPublicLiquidationData struct {
	Symbol      SymbolV5 `json:"s"`
	Side        Side     `json:"S"`
	Size        string   `json:"v"`
	Price       string   `json:"p"`
	UpdatedTime int64    `json:"T"`
}

// Key :
func (r *V5WebsocketPublicLiquidationResponse) Key() V5WebsocketPublicLiquidationParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicLiquidation.String() || len(arr) != 2 {
		return V5WebsocketPublicLiquidationParamKey{}
	}

	return V5WebsocketPublicLiquidationParamKey{
		Symbol: SymbolV5(arr[1]),
	}
}

// addParamLiquidationFunc :
func (v *V5WebsocketPublicService) addParamLiquidationFunc(key V5WebsocketPublicLiquidationParamKey, f func(*V5WebsocketPublicLiquidationResponse) error) error {
	if _, exist := v.paramLiquidationMap[key]; exist {
		return errors.New("already registered for this key")
	}
	v.paramLiquidationMap[key] = f
	return nil
}

// removeParamLiquidationFunc :
func (v *V5WebsocketPublicService) removeParamLiquidationFunc(key V5WebsocketPublicLiquidationParamKey) {
	delete(v.paramLiquidationMap, key)
}

// retrievePositionFunc :
func (v *V5WebsocketPublicService) retrieveLiquidationFunc(key V5WebsocketPublicLiquidationParamKey) (func(*V5WebsocketPublicLiquidationResponse) error, error) {
	f, exist := v.paramLiquidationMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
