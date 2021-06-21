package network

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/andrewarrow/cloutcli/keys"
	"github.com/btcsuite/btcd/btcec"
)

func SubmitTx(hexString string, priv *btcec.PrivateKey) string {
	jsonString := `{"TransactionHex": "%s"}`
	transactionBytes, _ := hex.DecodeString(hexString)
	first := sha256.Sum256(transactionBytes)
	transactionHash := sha256.Sum256(first[:])

	sig, _ := priv.Sign(transactionHash[:])
	signatureBytes := keys.SerializeToDer(sig)

	signatureLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(signatureLength, uint64(len(signatureBytes)))

	if len(transactionBytes) == 0 {
		return ""
	}

	buff := []byte{}
	buff = append(buff, transactionBytes[0:len(transactionBytes)-1]...)
	buff = append(buff, signatureLength[0])
	buff = append(buff, signatureBytes...)

	signedHex := fmt.Sprintf("%x", buff)

	send := fmt.Sprintf(jsonString, signedHex)
	jsonString = DoPost("api/v0/submit-transaction",
		[]byte(send))
	return jsonString
}
