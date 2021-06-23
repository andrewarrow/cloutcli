package main

import (
	"fmt"
	"sort"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/display"
	"github.com/andrewarrow/cloutcli/keys"
)

func HandleSell() {
	words := WarnAboutWords()
	if words == "" {
		return
	}
	//rate := cloutcli.GetRate()
	pub58, _ := keys.ComputeKeysFromSeed(words)
	me := cloutcli.Pub58ToUser(pub58)

	YouHODL := me.UsersYouHODL
	sort.SliceStable(YouHODL, func(i, j int) bool {
		vali := display.OneE9Float(YouHODL[i].BalanceNanos) * display.OneE9Float(YouHODL[i].ProfileEntryResponse.CoinPriceBitCloutNanos)
		valj := display.OneE9Float(YouHODL[j].BalanceNanos) * display.OneE9Float(YouHODL[j].ProfileEntryResponse.CoinPriceBitCloutNanos)
		return vali > valj
	})

	for _, thing := range YouHODL {
		if thing.HasPurchased == true {
			continue
		}
		val := display.OneE9Float(thing.BalanceNanos) * display.OneE9Float(thing.ProfileEntryResponse.CoinPriceBitCloutNanos)
		fmt.Printf("%20s %10s %10s %10s\n", thing.ProfileEntryResponse.Username,
			display.OneE9extra(thing.BalanceNanos),
			display.OneE9(thing.ProfileEntryResponse.CoinPriceBitCloutNanos),
			display.Float(val),
		)

	}

}
