package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/network"
)

func GlobalPosts() []Post {
	js := network.GetPostsStateless(DefaultPublicKey, false)
	var ps PostsStateless
	json.Unmarshal([]byte(js), &ps)

	return ps.PostsFound
}
func FollowingFeedPosts(username string) []Post {
	pub58 := UsernameToPub58(username)
	js := network.GetPostsStateless(pub58, true)
	var ps PostsStateless
	json.Unmarshal([]byte(js), &ps)

	return ps.PostsFound
}
