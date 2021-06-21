package network

import "fmt"

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
