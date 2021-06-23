package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/andrewarrow/cloutcli"
	"github.com/andrewarrow/cloutcli/display"
	"github.com/andrewarrow/cloutcli/keys"
)

func HandleSell() {
	words := WarnAboutWords()
	if words == "" {
		return
	}

	limit := argMap["limit"]
	execute := argMap["execute"]

	if limit == "" {
		fmt.Println("run with --limit=0.005")
		return
	}

	setLimit, _ := strconv.ParseFloat(limit, 64)

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
		if val > setLimit {
			continue
		}

		if execute != "" {
			fmt.Println("selling your", thing.ProfileEntryResponse.Username)
			ok := cloutcli.SubmitSellTransaction(words,
				thing.ProfileEntryResponse.PublicKeyBase58Check,
				thing.BalanceNanos)
			fmt.Println(ok)
		} else {
			fmt.Printf("%20s %10s %10s %10s\n", thing.ProfileEntryResponse.Username,
				display.OneE9extra(thing.BalanceNanos),
				display.OneE9(thing.ProfileEntryResponse.CoinPriceBitCloutNanos),
				fmt.Sprintf("%0.6f", val),
			)
		}

	}

}
