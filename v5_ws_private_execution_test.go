package bybit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dimkus/bybit/v2/testhelper"
)

func TestV5WebsocketPrivate_Execution(t *testing.T) {
	respBody := V5WebsocketPrivateExecutionResponse{
		ID:           "592324803b2785-26fa-4214-9963-bdd4727f07be",
		Topic:        "execution",
		CreationTime: 1672364174455,
		Data: []V5WebsocketPrivateExecutionData{
			{
				Category:        "linear",
				Symbol:          "XRPUSDT",
				ExecFee:         "0.005061",
				ExecID:          "7e2ae69c-4edf-5800-a352-893d52b446aa",
				ExecPrice:       "0.3374",
				ExecQty:         "25",
				ExecType:        ExecTypeV5Trade,
				ExecValue:       "8.435",
				IsMaker:         false,
				FeeRate:         "0.0006",
				TradeIv:         "",
				MarkIv:          "",
				BlockTradeID:    "",
				MarkPrice:       "0.3391",
				IndexPrice:      "",
				UnderlyingPrice: "",
				LeavesQty:       "0",
				OrderID:         "f6e324ff-99c2-4e89-9739-3086e47f9381",
				OrderLinkID:     "",
				OrderPrice:      "0.3207",
				OrderQty:        "25",
				OrderType:       OrderTypeUnknown,
				StopOrderType:   StopOrderTypeStopLoss,
				Side:            SideSell,
				ExecTime:        "1672364174443",
				IsLeverage:      "0",
				ClosedSize:      "",
				CrossSequence:   4688002127,
				CreateType:      CreateByClosing,
				ExecPnl:         "3.745",
			},
		},
	}
	bytesBody, err := json.Marshal(respBody)
	require.NoError(t, err)

	server, teardown := testhelper.NewWebsocketServer(
		testhelper.WithWebsocketHandlerOption(V5WebsocketPrivatePath, bytesBody),
	)
	defer teardown()

	wsClient := NewTestWebsocketClient().
		WithBaseURL(server.URL).
		WithAuth("test", "test")

	svc, err := wsClient.V5().Private()
	require.NoError(t, err)

	require.NoError(t, svc.Subscribe())

	{
		_, err := svc.SubscribeExecution(func(response V5WebsocketPrivateExecutionResponse) error {
			assert.Equal(t, respBody, response)
			return nil
		})
		require.NoError(t, err)
	}

	assert.NoError(t, svc.Run())
	assert.NoError(t, svc.Ping())
	assert.NoError(t, svc.Close())
}
