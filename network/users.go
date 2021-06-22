package network

import "fmt"

func GetUsersStateless(key string) string {
	jsonString := `{"PublicKeysBase58Check":["%s"],"SkipHodlings":false}`
	send := fmt.Sprintf(jsonString, key)
	jsonString = DoPost("api/v0/get-users-stateless",
		[]byte(send))
	return jsonString
}
