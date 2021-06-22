package cloutcli

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func GetNumFollowers(pub58, username string) int64 {
	js := network.GetFollowsStateless(pub58, username, "")
	var pktpe lib.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	return pktpe.NumFollowers
}

func LoopThruAllFollowing(pub58, username string, limit int) []lib.ProfileEntryResponse {
	last := ""
	js := network.GetFollowsStateless(pub58, username, last)
	var pktpe lib.PublicKeyToProfileEntry
	json.Unmarshal([]byte(js), &pktpe)
	NumFollowers := pktpe.NumFollowers
	total := map[string]bool{}
	bigList := []lib.ProfileEntryResponse{}
	fmt.Println("Getting all", pktpe.NumFollowers, "...")
	for {
		for key, v := range pktpe.PublicKeyToProfileEntry {
			last = key
			if total[v.Username] == false {
				total[v.Username] = true
				bigList = append(bigList, v)
			}
		}
		if len(bigList) >= limit && limit != 0 {
			break
		}
		if len(total) >= int(NumFollowers) {
			break
		}
		fmt.Println("got", len(bigList), "out of", NumFollowers)
		time.Sleep(time.Second * 1)
		js := network.GetFollowsStateless(pub58, username, last)
		json.Unmarshal([]byte(js), &pktpe)
	}
	return bigList
}
