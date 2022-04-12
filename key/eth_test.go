package key

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	btcSecp256k1 "github.com/LukeEuler/trx-go/key/btc_secp256k1"
)

func TestPubkeyToEthAddressBytes(t *testing.T) {
	raw := "0000000000000000000000000000000000000000000000000000000000000000"
	addrHex := "3f17f1962b36e491b30a40b2405849e597ba5fb5"

	bs, _ := hex.DecodeString(raw)

	private, _ := btcSecp256k1.PrivKeyFromBytes(btcSecp256k1.S256(), bs)
	privKey := private.ToECDSA()
	address := PubkeyToEthAddressBytes(privKey.PublicKey)

	assert.Equal(t, addrHex, hex.EncodeToString(address))
}
