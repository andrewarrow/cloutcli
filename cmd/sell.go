package main

import (
	"fmt"
	"sort"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/keys"
)

func HandleSell() {
	words := WarnAboutWords()
	if words == "" {
		return
	}
	pub58, _ := keys.ComputeKeysFromSeed(words)
	me := cloutcli.Pub58ToUser(pub58)

	YouHODL := me.UsersYouHODL
	sort.SliceStable(YouHODL, func(i, j int) bool {
		return YouHODL[i].BalanceNanos > YouHODL[j].BalanceNanos
	})

	for _, thing := range YouHODL {
		if thing.HasPurchased == true {
			continue
		}
		fmt.Printf("%20s %d %d\n", thing.ProfileEntryResponse.Username, thing.BalanceNanos, thing.ProfileEntryResponse.CoinPriceBitCloutNanos)
	}

}
