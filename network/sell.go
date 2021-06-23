package network

import "fmt"

func SubmitSellCoin(updater, creator string, sell, expected int64) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","CreatorPublicKeyBase58Check":"%s","OperationType":"sell","BitCloutToSellNanos":0,"CreatorCoinToSellNanos":%d,"BitCloutToAddNanos":0,"MinBitCloutExpectedNanos":%d,"MinCreatorCoinExpectedNanos":0,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, updater, creator, sell, expected)
	jsonString = DoPost("api/v0/buy-or-sell-creator-coin",
		[]byte(send))
	return jsonString
}
