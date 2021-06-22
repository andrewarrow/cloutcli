package network

import "fmt"

func GetMessagesStateless(pub58 string) string {
	jsonString := `{"PublicKeyBase58Check":"%s","FetchAfterPublicKeyBase58Check":"","NumToFetch":25,"HoldersOnly":false,"HoldingsOnly":false,"FollowersOnly":false,"FollowingOnly":false,"SortAlgorithm":"time"}`
	sendString := fmt.Sprintf(jsonString, pub58)
	jsonString = DoPost("api/v0/get-messages-stateless",
		[]byte(sendString))
	return jsonString
}
func SendMessage(sender, recipient, body string) string {
	jsonString := `{"SenderPublicKeyBase58Check":"%s","RecipientPublicKeyBase58Check":"%s","MessageText":"%s","MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, sender, recipient, body)
	jsonString = DoPost("api/v0/send-message-stateless",
		[]byte(send))
	return jsonString
}
