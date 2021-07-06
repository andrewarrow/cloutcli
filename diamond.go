package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func GiveDiamond(pub58, theirPub58, hash string) string {
	jsonString := network.SubmitDiamond("1", pub58, theirPub58, hash)

	var tx lib.TxReady
	json.Unmarshal([]byte(jsonString), &tx)
	return tx.TransactionHex
}
