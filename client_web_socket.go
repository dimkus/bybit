package bybit

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	// WebsocketBaseURL :
	WebsocketBaseURL = "wss://stream.bybit.com"
	// WebsocketBaseURL2 :
	WebsocketBaseURL2 = "wss://stream.bytick.com"
)

// WebSocketClient :
type WebSocketClient struct {
	debug  bool
	logger *log.Logger

	baseURL string
	key     string
	secret  string

	syncTimeDeltaNanoSeconds int64
}

func (c *WebSocketClient) debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Printf(format, v...)
	}
}

// NewWebsocketClient :
func NewWebsocketClient() *WebSocketClient {
	return &WebSocketClient{
		logger: newDefaultLogger(),

		baseURL: WebsocketBaseURL,
	}
}

// WithDebug :
func (c *WebSocketClient) WithDebug(debug bool) *WebSocketClient {
	c.debug = debug

	return c
}

// WithLogger :
func (c *WebSocketClient) WithLogger(logger *log.Logger) *WebSocketClient {
	c.debug = true
	c.logger = logger

	return c
}

// WithAuth :
func (c *WebSocketClient) WithAuth(key string, secret string) *WebSocketClient {
	c.key = key
	c.secret = secret

	return c
}

// WithBaseURL :
func (c *WebSocketClient) WithBaseURL(url string) *WebSocketClient {
	c.baseURL = url

	return c
}

// hasAuth : check has auth key and secret
func (c *WebSocketClient) hasAuth() bool {
	return c.key != "" && c.secret != ""
}

func (c *WebSocketClient) buildAuthParam() ([]byte, error) {
	if !c.hasAuth() {
		return nil, fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	expires := c.getTimestamp()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	s := hmac.New(sha256.New, []byte(c.secret))
	if _, err := s.Write([]byte(req)); err != nil {
		return nil, err
	}
	signature := hex.EncodeToString(s.Sum(nil))
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "auth",
		Args: []interface{}{c.key, expires, signature},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (c *WebSocketClient) getTimestamp() int64 {
	return (time.Now().UnixNano() - c.syncTimeDeltaNanoSeconds) / 1000000
}

// SetSyncTimeDeltaNs : set sync time delta in nanoseconds localTimeNs - remoteServerTimeNs
func (c *WebSocketClient) SetSyncTimeDeltaNs(timeDeltaNs int64) error {
	c.syncTimeDeltaNanoSeconds = timeDeltaNs
	return nil
}

func (c *WebSocketClient) UpdateSyncTimeDelta(
	remoteServerTimeNsRaw string,
	localTimestampNs int64,
) error {
	remoteServerTimeNs, err := strconv.ParseInt(remoteServerTimeNsRaw, 10, 64)
	if err != nil {
		return fmt.Errorf("parse server time: %w", err)
	}

	c.SetSyncTimeDeltaNs(localTimestampNs - remoteServerTimeNs)
	return nil
}

// WebsocketExecutor :
type WebsocketExecutor interface {
	Run() error
	Close() error
	Ping() error
}

// Start :
func (c *WebSocketClient) Start(ctx context.Context, executors []WebsocketExecutor) {
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			for _, executor := range executors {
				if err := executor.Run(); err != nil {
					if IsErrWebsocketClosed(err) {
						return
					}
					c.debugf("websocket executor error: %s", err)
					return
				}
			}
		}
	}()

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			for _, executor := range executors {
				if err := executor.Ping(); err != nil {
					return
				}
			}
		case <-ctx.Done():
			c.debugf("caught websocket interrupt signal")

			for _, executor := range executors {
				if err := executor.Close(); err != nil {
					return
				}
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
