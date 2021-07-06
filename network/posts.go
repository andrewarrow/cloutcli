package network

import (
	"fmt"
)

func GetPostsStateless(pub58 string, follow bool) string {
	jsonString := `{"GetPostsForGlobalWhitelist":%s,"GetPostsForFollowFeed":%s, "OrderBy":"newest", "ReaderPublicKeyBase58Check": "%s"}`

	withFollow := fmt.Sprintf(jsonString, "true", "false", pub58)
	if follow {
		withFollow = fmt.Sprintf(jsonString, "false", "true", pub58)
	}
	jsonString = DoPost("api/v0/get-posts-stateless",
		[]byte(withFollow))
	return jsonString
}

func SubmitPost(pub58, body, reply, imageURL string) string {
	jsonString := `{"UpdaterPublicKeyBase58Check":"%s","PostHashHexToModify":"","ParentStakeID":"%s","Title":"","BodyObj":{"Body":"%s","ImageURLs":[%s]},"RecloutedPostHashHex":"","PostExtraData":{},"Sub":"","IsHidden":false,"MinFeeRateNanosPerKB":1000}`
	send := fmt.Sprintf(jsonString, pub58, reply, body, imageURL)
	jsonString = DoPost("api/v0/submit-post",
		[]byte(send))
	return jsonString
}

func GetSinglePost(pub58, postHex string) string {
	jsonString := `{"PostHashHex":"%s","ReaderPublicKeyBase58Check":"%s","FetchParents":false,"CommentOffset":0,"CommentLimit":20,"AddGlobalFeedBool":false}`
	sendString := fmt.Sprintf(jsonString, postHex, pub58)
	jsonString = DoPost("api/v0/get-single-post",
		[]byte(sendString))
	return jsonString
}
