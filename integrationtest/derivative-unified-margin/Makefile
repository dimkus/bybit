BYBIT_TEST_UPDATED=false

test:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/dimkus/bybit/v2/integrationtest/derivative-unified-margin -v -tags=integrationtestderivativeunifiedmargin

test-spec:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/dimkus/bybit/v2/integrationtest/derivative-unified-margin -v -tags=integrationtestderivativeunifiedmargin -run ${BYBIT_TEST_METHOD}
