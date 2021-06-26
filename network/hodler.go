package network

import "fmt"

func GetHodlers(username string) string {
	jsonString := `{"PublicKeyBase58Check":"","Username":"%s","LastPublicKeyBase58Check":"","NumToFetch":100,"FetchHodlings":false,"FetchAll":false}`
	send := fmt.Sprintf(jsonString, username)
	jsonString = DoPost("api/v0/get-hodlers-for-public-key",
		[]byte(send))
	return jsonString
}
