package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/keys"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func SubmitSellTransaction(words, theirPub58 string, amount int64) string {
	pub58, priv := keys.ComputeKeysFromSeed(words)
	jsonString := network.SubmitSellCoin(pub58, theirPub58, amount, 0)
	var tx lib.TxReady
	json.Unmarshal([]byte(jsonString), &tx)

	jsonString = network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		return "ok"
	}
	return "error"
}
