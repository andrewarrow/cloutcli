package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func UsernamesOfHodlers(username string) []string {
	js := network.GetHodlers(username)
	var hw lib.HodlersWrap
	json.Unmarshal([]byte(js), &hw)

	items := []string{}
	for _, friend := range hw.Hodlers {
		username := friend.ProfileEntryResponse.Username
		if username != "" {
			items = append(items, username)
		}
	}
	return items
}
