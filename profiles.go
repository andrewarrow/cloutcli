package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func UsernameToPub58(s string) string {
	js := network.GetSingleProfile(s)
	var sp lib.SingleProfile
	json.Unmarshal([]byte(js), &sp)
	return sp.Profile.PublicKeyBase58Check
}

func Pub58ToUser(key string) lib.User {
	js := network.GetUsersStateless(key)
	var us lib.UsersStateless
	json.Unmarshal([]byte(js), &us)
	return us.UserList[0]
}
