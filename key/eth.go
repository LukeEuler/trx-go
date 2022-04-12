package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"hash"

	"golang.org/x/crypto/sha3"

	ethSecp256k1 "github.com/LukeEuler/trx-go/key/eth_secp256k1"
)

/*
github.com/ethereum/go-ethereum 这个库真的拉垮, 依赖容易冲突，故此处copy代码了(稍有改动)
github.com/ethereum/go-ethereum/crypto/secp256k1 -> github.com/LukeEuler/trx-go/key/eth_secp256k1
*/
func PubkeyToEthAddressBytes(p ecdsa.PublicKey) []byte {
	pubBytes := FromECDSAPub(&p)
	return Keccak256(pubBytes[1:])[12:]
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := NewKeccakState()

	for _, b := range data {
		d.Write(b)
	}
	_, _ = d.Read(b)
	return b
}

// KeccakState wraps sha3.state. In addition to the usual hash methods, it also supports
// Read to get a variable amount of data from the hash state. Read is faster than Sum
// because it doesn't copy the internal state, but also modifies the internal state.
type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

// NewKeccakState creates a new KeccakState
func NewKeccakState() KeccakState {
	return sha3.NewLegacyKeccak256().(KeccakState)
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(ethSecp256k1.S256(), pub.X, pub.Y)
}
