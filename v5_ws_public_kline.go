package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dimkus/websocket"
)

// SubscribeKline :
func (v *V5WebsocketPublicService) SubscribeKline(
	key V5WebsocketPublicKlineParamKey,
	f func(*V5WebsocketPublicKlineResponse) error,
) (func() error, error) {
	if err := v.addParamKlineFunc(key, f); err != nil {
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
		v.removeParamKlineFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicKlineParamKey :
type V5WebsocketPublicKlineParamKey struct {
	Interval Interval
	Symbol   SymbolV5
}

// Topic :
func (k *V5WebsocketPublicKlineParamKey) Topic() string {
	return fmt.Sprintf("%s.%s.%s", V5WebsocketPublicTopicKline, k.Interval, k.Symbol)
}

// V5WebsocketPublicKlineResponse :
// easyjson:json
type V5WebsocketPublicKlineResponse struct {
	Topic     string                       `json:"topic"`
	Type      string                       `json:"type"`
	Data      []V5WebsocketPublicKlineData `json:"data"`
	TimeStamp int64                        `json:"ts"`
}

// clear :
func (r *V5WebsocketPublicKlineResponse) clear() {
	r.Topic = ""
	r.Type = ""
	r.TimeStamp = 0
	if cap(r.Data) > 40 {
		r.Data = nil
	} else {
		for i := range r.Data {
			r.Data[i] = V5WebsocketPublicKlineData{}
		}
		r.Data = r.Data[:0]
	}
}

// V5WebsocketPublicKlineData :
// easyjson:json
type V5WebsocketPublicKlineData struct {
	Interval  Interval `json:"interval"`
	Open      string   `json:"open"`
	Close     string   `json:"close"`
	High      string   `json:"high"`
	Low       string   `json:"low"`
	Volume    string   `json:"volume"`
	Turnover  string   `json:"turnover"`
	Start     int64    `json:"start"`
	End       int64    `json:"end"`
	Timestamp int64    `json:"timestamp"`
	Confirm   bool     `json:"confirm"`
}

// clear :
func (r *V5WebsocketPublicKlineData) clear() {
	r.Interval = ""
	r.Open = ""
	r.Close = ""
	r.High = ""
	r.Low = ""
	r.Volume = ""
	r.Turnover = ""
	r.Start = 0
	r.End = 0
	r.Timestamp = 0
	r.Confirm = false
}

// Key :
func (r *V5WebsocketPublicKlineResponse) Key() V5WebsocketPublicKlineParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicKline.String() || len(arr) != 3 {
		return V5WebsocketPublicKlineParamKey{}
	}

	return V5WebsocketPublicKlineParamKey{
		Interval: Interval(arr[1]),
		Symbol:   SymbolV5(arr[2]),
	}
}

// addParamKlineFunc :
func (v *V5WebsocketPublicService) addParamKlineFunc(key V5WebsocketPublicKlineParamKey, f func(*V5WebsocketPublicKlineResponse) error) error {
	if _, exist := v.paramKlineMap[key]; exist {
		return errors.New("already registered for this key")
	}
	v.paramKlineMap[key] = f
	return nil
}

// removeParamTradeFunc :
func (v *V5WebsocketPublicService) removeParamKlineFunc(key V5WebsocketPublicKlineParamKey) {
	delete(v.paramKlineMap, key)
}

// retrievePositionFunc :
func (v *V5WebsocketPublicService) retrieveKlineFunc(key V5WebsocketPublicKlineParamKey) (func(*V5WebsocketPublicKlineResponse) error, error) {
	f, exist := v.paramKlineMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
