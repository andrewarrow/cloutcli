package cloutcli

import (
	"encoding/json"

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
