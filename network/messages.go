package network

import "fmt"

func GetMessagesStateless(pub58 string) string {
	jsonString := `{"PublicKeyBase58Check":"%s","FetchAfterPublicKeyBase58Check":"","NumToFetch":25,"HoldersOnly":false,"HoldingsOnly":false,"FollowersOnly":false,"FollowingOnly":false,"SortAlgorithm":"time"}`
	sendString := fmt.Sprintf(jsonString, pub58)
	jsonString = DoPost("api/v0/get-messages-stateless",
		[]byte(sendString))
	return jsonString
}
