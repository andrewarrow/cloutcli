package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func GlobalPosts() []lib.Post {
	js := network.GetPostsStateless(DefaultPublicKey, false)
	var ps lib.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	return ps.PostsFound
}
func FollowingFeedPosts(username string) []lib.Post {
	pub58 := UsernameToPub58(username)
	js := network.GetPostsStateless(pub58, true)
	var ps lib.PostsStateless
	json.Unmarshal([]byte(js), &ps)

	return ps.PostsFound
}
