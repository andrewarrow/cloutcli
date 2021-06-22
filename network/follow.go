package network

import "fmt"

func GetFollowsStateless(pub58, username, last string) string {
	jsonString := `{"Username":"%s","PublicKeyBase58Check":"%s","GetEntriesFollowingUsername":%s,"LastPublicKeyBase58Check":"%s","NumToFetch":50}`

	withDirection := fmt.Sprintf(jsonString, username, pub58, "false", last)
	if username != "" {
		withDirection = fmt.Sprintf(jsonString, username, pub58, "true", last)
	}

	jsonString = DoPost("api/v0/get-follows-stateless",
		[]byte(withDirection))
	return jsonString
}
