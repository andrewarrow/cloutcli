package keys

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
)

func SerializeToDer(sig *btcec.Signature) []byte {
	r := sig.R.Bytes()
	s := sig.S.Bytes()

	rBuff := []string{}
	sBuff := []string{}

	for _, b := range r {
		rBuff = append(rBuff, fmt.Sprintf("%d", b))
	}
	for _, b := range s {
		sBuff = append(sBuff, fmt.Sprintf("%d", b))
	}

	payload := RunV8(strings.Join(rBuff, ","),
		strings.Join(sBuff, ","))
	res := []byte{}
	for _, b := range strings.Split(payload, ",") {
		thing, _ := strconv.Atoi(b)
		res = append(res, byte(thing))
	}
	return res
}
