package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/network"
)

func UsernameToPub58(s string) string {
	js := network.GetSingleProfile(s)
	var sp SingleProfile
	json.Unmarshal([]byte(js), &sp)
	return sp.Profile.PublicKeyBase58Check
}
