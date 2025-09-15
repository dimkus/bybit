package bybit

import "github.com/dimkus/websocket"

// IsErrWebsocketClosed :
func IsErrWebsocketClosed(err error) bool {
	return websocket.IsCloseError(err, websocket.CloseNormalClosure)
}
