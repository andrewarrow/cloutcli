package network

import "fmt"

func SubmitDiamond(level, sender, receiver, post string) string {
	jsonString := `{"SenderPublicKeyBase58Check":"%s","ReceiverPublicKeyBase58Check":"%s","DiamondPostHashHex":"%s","DiamondLevel":%s,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, sender, receiver, post, level)
	jsonString = DoPost("api/v0/send-diamonds",
		[]byte(send))
	return jsonString
}
