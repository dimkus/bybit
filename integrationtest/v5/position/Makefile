BYBIT_TEST_UPDATED=false

test:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/dimkus/bybit/v2/integrationtest/v5/position -v -tags=integrationtestv5position

test-spec:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/dimkus/bybit/v2/integrationtest/v5/position -v -tags=integrationtestv5position -run ${BYBIT_TEST_METHOD}

