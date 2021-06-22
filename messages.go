package cloutcli

import (
	"encoding/json"

	"github.com/andrewarrow/cloutcli/keys"
	"github.com/andrewarrow/cloutcli/lib"
	"github.com/andrewarrow/cloutcli/network"
)

func MessageInbox(username string) lib.MessageList {
	pub58 := UsernameToPub58(username)
	js := network.GetMessagesStateless(pub58)
	var list lib.MessageList
	json.Unmarshal([]byte(js), &list)

	return list
}
func SendMessage(words, to, body string) string {
	pub58, priv := keys.ComputeKeysFromSeed(words)
	recipient := UsernameToPub58(to)
	jsonString := network.SendMessage(pub58, recipient, body)
	var tx lib.TxReady
	json.Unmarshal([]byte(jsonString), &tx)

	jsonString = network.SubmitTx(tx.TransactionHex, priv)
	if jsonString != "" {
		return "ok"
	}
	return "error"
}
